package analysis

import (
  "btce"
  "fmt"
)

type Status struct {
  Time     int32   // BTC-E server unix time in seconds
  Price    float32 // Price of last trade
  Analysis Analysis
}

// This is the entry point that happens on a frequent interval
// that pulls the newest information about the market, analyzes it,
// ands sends it off to be looked at (and possibly acted on) by
// the troll.

func RecordMarketStatus() (status Status) {
  // Get the last price for which BTC was traded
  lastTrade := btce.GetLastTrade()
  status.Time = lastTrade.Date
  status.Price = lastTrade.Price

  // Pre-process analysis on this price
  status.Analysis = Analyze(status)

  // Save it in Mongo
  err := coll.Insert(&status)
  if err != nil {
    fmt.Println(fmt.Sprintf("Error: %s", err))
  }

  return status
}

