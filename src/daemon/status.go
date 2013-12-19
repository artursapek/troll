package daemon

import (
  "data"
  "btce"
  "fmt"
)

type StatusDaemon struct{}

func getStatus() string {
  c := data.GetCollection("prices")
  var tickers []btce.Ticker
  c.Find(nil).Limit(1).Sort("-server_time").All(&tickers)
  lastTicker := tickers[0]
  time := lastTicker.Server_time
  buy := lastTicker.Buy
  sell := lastTicker.Sell
  return fmt.Sprintf("\r%d  %f  %f", time, buy, sell)
}

func (daemon StatusDaemon) Setup() {
  header := "time        buy         sell"
  fmt.Println(header)
}

func (daemon StatusDaemon) Perform() {
  fmt.Printf(getStatus())
}
