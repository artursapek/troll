package btce

import (
  "encoding/json"
)

type PublicTrade struct {
  Date   int64
  Price  float32
}

type PublicTrades []PublicTrade

func decodePublicTrades(input []byte) PublicTrades {
  var trades PublicTrades
  err := json.Unmarshal(input, &trades)
  if err != nil {
    panic(err)
  }
  return trades
}

func GetPublicTrades() PublicTrades {
  response := PublicApiRequest("trades")
  return decodePublicTrades(response)
}

func GetLastTrade() PublicTrade {
  return GetPublicTrades()[0]
}

