package daemon

import (
  "data"
  "btce"
  "fmt"
)

type StatusDaemon struct{}

func (daemon StatusDaemon) Perform() {
  c := data.GetCollection("prices")
  var lastPrice []btce.Ticker
  c.Find(nil).Limit(1).Sort("-server_time").All(&lastPrice)
  fmt.Println(lastPrice[0].Buy)
}
