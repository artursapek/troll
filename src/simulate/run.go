package simulate

import (
  "fmt"
  "data"
  "btce"
  "state"
  "analysis"
)

// Mongo db holding test data
const testDB string = "test_prices"
const amtDocs int = 70583

var c = data.GetCollection(testDB)
var cc = data.GetCollection("test_prices_analyzed")

var firstTrade = btce.Trade{
  Pair: "btc_usd",
  Type: "sell",
  Amount: 0.5,
  Rate: 206.0,
  Order_id: 1,
  Timestamp: 1381776000,
}

var testState = state.TrollState{
  LastTrade: firstTrade,
}

var i = 5001;
var statuses []analysis.Status
var err = c.Find(nil).All(&statuses)

func Iterate() {
  status := statuses[i]
  status = analysis.Analyze(status)
  fmt.Println(status)
  cc.Insert(&status)
  i += 1
  if i < 10000 {
    Iterate()
  }
}

