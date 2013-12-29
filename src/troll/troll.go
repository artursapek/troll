package troll

import (
  "btce"
//  "market"
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
  //status := market.RecordPrice()
  //troll.Decide(status)
  return time.Duration(15)
}


