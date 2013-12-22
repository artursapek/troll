package state

import(
  "fmt"
  "btce"
)

type TrollState struct {
  LastTrade btce.Trade
}

var State TrollState

func GetState() {
  lastTrade := btce.LastTrade()
  fmt.Println(lastTrade)
}


