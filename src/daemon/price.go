package daemon

import (
  "btce"
  "data"
  "fmt"
)

type PriceDaemon struct {}

func (daemon PriceDaemon) Perform() {
  ticker := btce.GetTicker()
  // If ticker isn't empty (meaning there was a decode error)
  if ticker != (btce.Ticker{}) {
    c := data.GetCollection("prices")
    err := c.Insert(&ticker)
    if err != nil {
      panic(err)
    } else {
      fmt.Println("Prices saved")
    }
  } else {
    fmt.Println("Ignoring empty ticker")
  }
}


