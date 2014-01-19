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
  json.Unmarshal(input, &trades)
  return trades
}

func GetPublicTrades() PublicTrades {
  response := PublicApiRequest("trades")
  return decodePublicTrades(response)
}

func GetLastTrade() PublicTrade {
  trades := GetPublicTrades()
  if len(trades) > 0 {
    return trades[0]
  } else {
    return PublicTrade{}
  }
}

