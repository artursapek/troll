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

// Mongo db holding test data
const testDB string = "test_prices"
const amtDocs int = 70583

var testCollection = data.GetStatusCollection()

// Set up test state
var funds = troll.FundsStatus{
  BTC: 0.0,
  USD: 103.25,
}

var self = troll.Troll{ 
  Funds: funds,
  LastTrade: btce.OwnTrade{
    Pair: "btc_usd",
    Type: "sell",
    Amount: 0.5,
    Rate: 206.5,
    Timestamp: 1385577480,
  },
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

  fmt.Println(self.Funds)

  var statuses []analysis.MarketStatus
  testCollection.Find(nil).Skip(skip).Limit(limit).All(&statuses)

  for i := 0; i < limit; i ++ {
    status := statuses[i]
    status = analysis.Analyze(status)
    self = self.Decide(status)
  }
}

