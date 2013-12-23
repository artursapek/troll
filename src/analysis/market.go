package analysis

import (
  "btce"
  "fmt"
  "data"
  "time"
  "labix.org/v2/mgo/bson"
)

var coll = data.GetCollection("statuses")

type MarketStatus struct {
  ServerTime int32 // BTC-E server unix time in seconds
  LocalTime  int32 // Local time
  Price    float32 // Price of last trade
  Analysis Analysis
}

// This is the entry point that happens on a frequent interval
// that pulls the newest information about the market, analyzes it,
// ands sends it off to be looked at (and possibly acted on) by
// the troll.

func RecordMarketStatus() (status MarketStatus) {

  // Get the last price for which BTC was traded
  lastTrade := btce.GetLastTrade()

  count, _ := coll.Find(bson.M{ "time": lastTrade.Date }).Count()
  if count > 0 {
    // If there are no new trades since last time, return an empty status
    fmt.Println("No new trades; skipping")
    return status
  }

  status.ServerTime = lastTrade.Date
  status.Price = lastTrade.Price
  status.LocalTime = int32(time.Now().Unix())

  // Pre-process analysis on this price
  status = Analyze(status)

  // Save it in Mongo
  err := coll.Insert(&status)
  if err != nil {
    fmt.Println(fmt.Sprintf("Error: %s", err))
  }

  return status
}

