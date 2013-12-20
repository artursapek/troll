package simulate

import (
//  "fmt"
  "data"
  "daemon"
  "btce"
  "state"
)

// Mongo db holding test data
const testDB string = "test_prices"

const amtDocs int = 70583

func Iterate() {
  state.GetState()

  c := data.GetCollection(testDB)
  var ticker btce.Ticker
  for skip := 0; skip < 1000; skip ++ {
    var tickers []btce.Ticker
    c.Find(nil).Limit(1).Skip(skip).All(&tickers)
    ticker = tickers[0]
    //fmt.Printf(".")
    Run(ticker)
  }
}

func Run(ticker btce.Ticker) {
  daemon.Tick(ticker)
}

