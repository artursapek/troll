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
  c := data.GetCollection("prices")
  err := c.Insert(&ticker)
  if err != nil {
    panic(err)
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
