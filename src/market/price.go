package market

import (
  "btce"
  "fmt"
  "data"
  "time"
  "log"
  "labix.org/v2/mgo/bson"
)

type MarketPrice struct {
  Time struct {
    Server, Local int64
  }
  Price float32 // Last trade conducted
}

func alreadyRecorded(time int64) bool {
  count, _ := data.Prices.Find(bson.M{ "time": time }).Count()
  return count > 0
}

func getPricesBetween (start, end int64) (prices []MarketPrice) {
  // Find all statuses between the two dates, excluding the given status
  query := bson.M{ "time.server": bson.M{ "$gt": start, "$lte": end }}
  err := data.Prices.Find(query).Sort("-time.server").All(&prices)
  if err != nil {
    fmt.Println(err) // Log it, and return empty
  }
  return prices
}

func getFirstPrice() (price MarketPrice) {
  var prices []MarketPrice
  data.Prices.Find(nil).Limit(1).Sort("time.server").All(&prices)
  return prices[0]
}

func RecordPrice() (price MarketPrice) {
  // Get the last price for which BTC was traded
  lastTrade := btce.GetLastTrade()

  // Make sure we haven't already recorded this trade
  if alreadyRecorded(lastTrade.Date) {
    // If we have, return an empty price
    fmt.Println("No new trades; skipping")
    return price
  }

  price.Time.Server = lastTrade.Date
  price.Price = lastTrade.Price
  price.Time.Local = int64(time.Now().Unix())

  // Persist to Mongo and we're done
  err := data.Prices.Insert(&price)
  if err != nil {
    log.Println("ERROR FROM MONGO")
    log.Println(err)
  }

  // Summarize 2-hour interval if it's time to
  ProcessPrice(price)

  return price
}

func ProcessPrice(price MarketPrice) {
  lastClose := lastIntervalCloseTime()
  if price.Time.Local - lastClose >= INTERVAL_PERIOD {
    // If it's been at least two hours since the last interval,
    // let's record the newest.
    RecordInterval(lastClose)
  }
}

