package troll

import (
  "btce"
  "analysis"
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

func (troll Troll) Setup() {}

func (troll Troll) Perform() time.Duration {
  // Record the current market price and analyze it
  status := analysis.RecordMarketStatus()
  //troll.Decide(status)
  if troll.Excited(status) {
    // Check more often when something is happening
    return time.Duration(15)
  } else {
    return time.Duration(30)
  }
}


