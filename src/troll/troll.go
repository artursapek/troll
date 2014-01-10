package troll

import (
  "btce"
  "market"
  "time"
  "data"
  "fmt"
)

const CLR_WHITE  = "\x1b[37;1m"
const CLR_GREY   = "\x1b[30;1m"
const CLR_GREEN  = "\x1b[32;1m"
const CLR_YELLOW = "\x1b[33;1m"
const CLR_RED    = "\x1b[31;1m"

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
    market.RecordInterval(lastClose)
    //self.Decide(interval)
  }

  return time.Duration(30)
}

func BuildIntervals() {
  data.Intervals.DropCollection()

  var prices []market.MarketPrice
  data.Prices.Find(nil).All(&prices)

  for _, price := range prices {

    lastClose, isDue := market.CheckIfNewIntervalIsDue(price.Time.Local)

    if isDue {
      market.RecordInterval(lastClose)
      fmt.Printf(".")
    }
  }

}


