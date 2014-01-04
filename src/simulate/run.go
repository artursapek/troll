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

func Simulate() {
  data.Trades.DropCollection()
  fmt.Println("Simulating...")
  var skip, limit int
  if len(os.Args) < 3 {
    skip = 0
    limit = 999999999999
  } else {
    skip, _ = strconv.Atoi(os.Args[2])
    limit, _ = strconv.Atoi(os.Args[3])
  }


  var intervals []market.MarketInterval
  data.Intervals.Find(nil).Skip(skip).Limit(limit).Sort("time.close").All(&intervals)

  //var prices []market.MarketPrice
  //data.Prices.Find(nil).Skip(skip).Limit(limit).All(&prices)

  amt := len(intervals)

  self := MakeTrollFromStatus(intervals[0].CandleStick.Open, intervals[0].Time.Open)

  for i := 0; i < amt; i ++ {
    interval := intervals[i]
    interval = market.AnalyzeInterval(interval)

    market.PersistUpdatedInterval(interval)

  /*
    fmt.Printf("%d,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f\n", interval.Time.Close,
               interval.CandleStick.Open,
               interval.CandleStick.Close,
               interval.CandleStick.High,
               interval.CandleStick.Low,
               interval.Ichimoku.TenkenSen,
               interval.Ichimoku.KijunSen,
               interval.Ichimoku.SenkouSpanA,
               interval.Ichimoku.SenkouSpanB,
               interval.Ichimoku.ChikouSpan,
               interval.SAR.Value)
               */
    self = self.Decide(interval)
  }
  /*
  for i := 0; i < amt; i ++ {
    price := prices[i]
    //now := price.Time.Local

    lastClose, isDue := market.CheckIfNewIntervalIsDue(price.Time.Local)

    if isDue {
      interval := market.RecordInterval(lastClose)
      self = self.Decide(interval)
    }
  }
  */
}

