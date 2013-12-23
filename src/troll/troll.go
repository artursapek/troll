package troll

import (
  "btce"
  "analysis"
)

type Troll struct {
  Funds FundsStatus
  LastTrade btce.OwnTrade
}

func CreateSyncedTroll() Troll {
  // Sync up with btc-e, make a troll
  funds := GetFundsStatus() 
  return Troll{
    Funds: funds,
    LastTrade: btce.LastTradeMade(),
  }
}

func (troll Troll) Setup() {}

func (troll Troll) Perform() {
  // Record the current market price and analyze it
  status := analysis.RecordMarketStatus()
  troll.Decide(status)
}


