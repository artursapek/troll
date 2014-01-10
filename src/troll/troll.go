package troll

import (
  "fmt"
  "btce"
  "market"
  "time"
  "data"
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
  lastInterval := market.PastNIntervals(1)[0]

  if isDue {
    market.RecordInterval(lastClose, lastInterval.CandleStick.Close)
    //self.Decide(interval)
  }

  return time.Duration(30)
}

func BuildIntervals() {
  data.Intervals.DropCollection()

  var lastCloseTime int64
  var lastClosePrice float32

  var lastInterval market.MarketInterval

  var prices      []market.MarketPrice
  var pricesStack []market.MarketPrice

  data.Prices.Find(nil).Sort("time.local").All(&prices)

  for i, price := range prices {

    if i == 0 {
      lastCloseTime = price.Time.Local
      lastClosePrice = price.Price
    }

    if price.Time.Local - lastCloseTime >= market.INTERVAL_PERIOD {
      fmt.Println("found new")
      // Time for a new interval!
      lastInterval = market.RecordInterval(lastCloseTime, lastClosePrice)
      lastCloseTime = lastInterval.Time.Close
      lastClosePrice = lastInterval.CandleStick.Close
      pricesStack = []market.MarketPrice{}
    } else {
      pricesStack = append(pricesStack, price)
    }
  }

}


