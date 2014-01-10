package simulate

import (
  "fmt"
  "os"
  "data"
  "btce"
  "market"
  "troll"
  "strconv"
)

// Start at any point in the simulation
func MakeTrollFromStatus(price float32, time int64) troll.Troll {
  funds := troll.FundsStatus{
    BTC: 0,
    USD: price,
  }
  return troll.Troll{ 
    Funds: funds,
    LastTrade: btce.OwnTrade{
      Pair: "btc_usd",
      Type: "sell",
      Amount: 1.0,
      Rate: price,
      Timestamp: time,
    },
  }
}

var skip, limit int

func parseSkipLimit() {
  if len(os.Args) < 3 {
    skip = 0
    limit = 999999999999
  } else {
    skip, _ = strconv.Atoi(os.Args[2])
    limit, _ = strconv.Atoi(os.Args[3])
  }
}



func Trade() {
  parseSkipLimit()
  data.Trades.DropCollection()
  fmt.Println("Simulating...")

  var intervals []market.MarketInterval
  data.Intervals.Find(nil).Skip(skip).Limit(limit).Sort("time.close").All(&intervals)

  amt := len(intervals)

  self := MakeTrollFromStatus(intervals[0].CandleStick.Open, intervals[0].Time.Open)

  for i := 0; i < amt; i ++ {
    interval := intervals[i]
    interval = market.AnalyzeInterval(interval)

    market.PersistUpdatedInterval(interval)
    self = self.Decide(interval)
  }
}


func BuildIntervals() {
  parseSkipLimit()
  data.Intervals.DropCollection()

  var prices []market.MarketPrice
  data.Prices.Find(nil).Skip(skip).Limit(limit).All(&prices)

  for _, price := range prices {

    lastClose, isDue := market.CheckIfNewIntervalIsDue(price.Time.Local)
    lastInterval := market.PastNIntervals(1)[0]

    if isDue {
      market.RecordInterval(lastClose, lastInterval.CandleStick.Close)
    }
  }
}

