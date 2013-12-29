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
func MakeTrollFromStatus(status market.MarketPrice) troll.Troll {
  funds := troll.FundsStatus{
    BTC: 1.0,
    USD: 0,
  }
  return troll.Troll{ 
    Funds: funds,
    LastTrade: btce.OwnTrade{
      Pair: "btc_usd",
      Type: "buy",
      Amount: 1.0,
      Rate: status.Price,
      Timestamp: status.Time.Server,
    },
  }

}

func Simulate() {
  var skip, limit int
  if len(os.Args) < 3 {
    skip = 0
    limit = 10000
  } else {
    skip, _ = strconv.Atoi(os.Args[2])
    limit, _ = strconv.Atoi(os.Args[3])
  }

  var statuses []market.MarketPrice
  data.Prices.Find(nil).Skip(skip).Limit(limit).All(&statuses)

  self := MakeTrollFromStatus(statuses[0])

  fmt.Println(self.Funds)
  /*
  for i := 0; i < limit; i ++ {
    status := statuses[i]
    status = market.Analyze(status)
    self = self.Decide(status)
  }
  */
}

