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
  query := bson.M{ "servertime": bson.M{ "$gte": start, "$lt": end }}
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

  return price
}

func ProcessPrice(price MarketPrice) {
  if NewIntervalHasClosed(price.Time.Local) {
    
  }
}

