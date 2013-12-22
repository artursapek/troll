package daemon

import (
  "fmt"
  "analysis"
)

const BUY string = "buy"
const SELL string = "sell"

type TrollDaemon struct{}

func (daemon TrollDaemon) Perform() {
  Update()
}

func (daemon TrollDaemon) Setup() {
}

func Update() {
  status := analysis.RecordMarketStatus()
  // Act on it! Buy or wait, sell or hold
  fmt.Println(status)
}

