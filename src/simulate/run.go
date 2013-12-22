package simulate

import (
  "data"
  "btce"
  "state"
)

// Mongo db holding test data
const testDB string = "test_prices"
const amtDocs int = 70583


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


func Iterate() {
  //state.GetState()
  skip := 0

  c := data.GetCollection(testDB)
  var tickers []btce.Ticker
  c.Find(nil).Skip(skip).All(&tickers)

  for i := 0; i < len(tickers); i ++ {
    //ticker := tickers[i]
  }
}

