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
    Server, Local int32
  }
  Price float32 // Last trade conducted
}

func alreadyRecorded(time int32) bool {
  count, _ := data.Prices.Find(bson.M{ "time": time }).Count()
  return count > 0
}

func pricesBetween (start, end int32) (prices []MarketPrice) {
  // Find all statuses between the two dates, excluding the given status
  query := bson.M{ "servertime": bson.M{ "$gte": start, "$lt": end }}
  err := data.Prices.Find(query).Sort("-servertime").All(&prices)
  if err != nil {
    fmt.Println(err) // Log it, and return empty
  }
  return prices
}

func RecordPrice() (status MarketPrice) {
  // Get the last price for which BTC was traded
  lastTrade := btce.GetLastTrade()

  // Make sure we haven't already recorded this trade
  if alreadyRecorded(lastTrade.Date) {
    // If we have, return an empty status
    fmt.Println("No new trades; skipping")
    return status
  }

  status.Time.Server = lastTrade.Date
  status.Price = lastTrade.Price
  status.Time.Local = int32(time.Now().Unix())

  // Persist to Mongo and we're done
  err := data.Prices.Insert(&status)
  if err != nil {
    log.Println("ERROR FROM MONGO")
    log.Println(err)
  }

  return status
}

