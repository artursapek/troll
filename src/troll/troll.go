package troll

import (
  "btce"
  "market"
  "time"
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

  lastClose, isDue := market.CheckIfNewIntervalIsDue(price.Time.Local)

  if isDue {
    interval := market.RecordInterval(lastClose)
    self.Decide(interval)
  }

  return time.Duration(15)
}


