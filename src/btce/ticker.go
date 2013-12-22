package btce

import (
  "fmt"
  "encoding/json"
)

type Ticker struct {
  Buy float32
  Sell float32
  Price float32
  Server_time int32
}

// The JSON from their api returns { ticker: { ... everything }}
// It's stupid
type TickerUnpack struct {
  Ticker Ticker
}

func decodeTicker(input []byte) Ticker {
  var ticker TickerUnpack
  err := json.Unmarshal(input, &ticker)
  if err != nil {
    fmt.Println(err)
  }
  return ticker.Ticker
}

func GetTicker() Ticker {
  response := PublicApiRequest("ticker")
  return decodeTicker(response)
}

