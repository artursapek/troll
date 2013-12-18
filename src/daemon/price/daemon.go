package price

import (
  "btce"
  "data"
  "time"
  "fmt"
)

const frequency = 30 // seconds

func perform() {
  fmt.Println("Running...")
  ticker := btce.GetTicker()
  if ticker != (btce.Ticker{}) {
    c := data.GetCollection("prices")
    err := c.Insert(&ticker)
    if err != nil {
      panic(err)
    }
  } else {
    fmt.Println("Skipping empty ticker")
  }
}

func Run() {
  running := true
  for running {
    perform()
    // Sleep for a minute
    time.Sleep(frequency * 1000 * time.Millisecond)
  }
}
