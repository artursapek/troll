package simulate

import (
  "fmt"
  "os"
  "data"
  "btce"
  "analysis"
  "troll"
  "strconv"
)

var testCollection = data.GetStatusCollection()

// Start at any point in the simulation
func MakeTrollFromStatus(status analysis.MarketStatus) troll.Troll {
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
      Timestamp: status.ServerTime,
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

  var statuses []analysis.MarketStatus
  testCollection.Find(nil).Skip(skip).Limit(limit).All(&statuses)

  self := MakeTrollFromStatus(statuses[0])

  fmt.Println(self.Funds)

  for i := 0; i < limit; i ++ {
    status := statuses[i]
    status = analysis.Analyze(status)
    self = self.Decide(status)
  }
}

