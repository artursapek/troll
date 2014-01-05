package monitor

import (
  "fmt"
  "net/http"
  "data"
  "market"
  "encoding/json"
  "btce"
)

func intervalsHandler(rw http.ResponseWriter, req *http.Request) {
  var intervals market.MarketIntervals
  data.Intervals.Find(nil).All(&intervals)
  body, err := json.Marshal(intervals)
  if err != nil {
    panic(err)
  } else {
    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(rw, string(body))
  }
}

func tradesHandler(rw http.ResponseWriter, req *http.Request) {
  var trades []btce.OwnTrade
  data.Trades.Find(nil).All(&trades)
  body, _ := json.Marshal(trades)

  rw.Header().Set("Access-Control-Allow-Origin", "*")
  rw.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(rw, string(body))
}

func StartServer() {
  http.HandleFunc("/prices.json", intervalsHandler)
  http.HandleFunc("/trades.json", tradesHandler)
  http.ListenAndServe(":8001", nil)
}
