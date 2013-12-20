package state

import(
  "fmt"
  "btce"
)

type TrollState struct {
  WaitingTo string // "buy"/"sell"
}

func GetState() {
  lastTrade := btce.LastTrade()
  fmt.Println(lastTrade)
}


