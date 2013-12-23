package simulate

import (
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

var testCollection = data.GetCollection(testDB)
var cc = data.GetCollection("test_prices_analyzed")

// Set up test state
var funds = troll.FundsStatus{
  BTC: 0.0,
  USD: 102.9,
}

var self = troll.Troll{ 
  Funds: funds,
  LastTrade: btce.OwnTrade{
    Pair: "btc_usd",
    Type: "sell",
    Amount: 0.5,
    Rate: 206.0,
    Timestamp: 1381776000,
  },
}


func Simulate() {
  skip, _ := strconv.Atoi(os.Args[2])
  limit, _ := strconv.Atoi(os.Args[3])

  var statuses []analysis.MarketStatus
  testCollection.Find(nil).Skip(skip).Limit(limit).All(&statuses)

  for i := 0; i < limit; i ++ {
    status := statuses[i]
    status = analysis.Analyze(status)
    self.Decide(status)
  }
}

