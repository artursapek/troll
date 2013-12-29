package troll

import (
  "btce"
  "data"
  "fmt"
)

func GetTicker() btce.Ticker {
  ticker := btce.GetTicker()
  // If ticker is empty there was a decode error
  if ticker != (btce.Ticker{}) {
    err := data.Prices.Insert(&ticker)
    if err != nil {
      panic(err)
    } else {
      fmt.Println("Prices saved")
    }
  } else {
    fmt.Println("Ignoring empty ticker")
  }
  return ticker
}


