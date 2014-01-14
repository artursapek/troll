package troll

import (
  "fmt"
  "btce"
  "market"
  "time"
  "data"
)

type Troll struct {
  Live bool
  Funds FundsStatus
  LastTrade btce.OwnTrade
}

func CreateSyncedTroll() Troll {
  // Sync up with btc-e, make a troll
  funds := GetFundsStatus() 
  return Troll{
    Live: true,
    Funds: funds,
    LastTrade: btce.LastTradeMade(),
  }
}

func (self Troll) Setup() { /* noop */ }

func (self Troll) Perform() time.Duration {
  // Record the current market price
  price := market.RecordPrice()
  self.ProcessPrice(price)
  return time.Duration(30)
}

// Cache the last interval
var lastInterval market.MarketInterval

func (self Troll) ProcessPrice(price market.MarketPrice) {
  if (lastInterval == market.MarketInterval{}) {
    // Query and cache the last interval's close time if we haven't
    lastInterval = market.LastInterval()
  }

  if price.Time.Local - lastInterval.Time.Close >= market.INTERVAL_PERIOD {
    // Record a new interval if 2 hours has passed and
    // update the cache
    lastInterval = market.RecordIntervalSucceeding(lastInterval)
  }
}

func RebuildIntervals() {
  // Rebuild all intervals from the beginning
  fmt.Println("Dropping intervals...")
  data.Intervals.DropCollection()

  self := Troll{}

  fmt.Println("Getting prices...")
  var prices []market.MarketPrice

  data.Prices.Find(nil).Sort("time.local").All(&prices)

  for i, price := range prices {
    if i == 0 {
      lastInterval.Time.Close = price.Time.Local
    }

    self.ProcessPrice(price)
  }
}

