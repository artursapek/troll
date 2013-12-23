package btce

import (
  "fmt"
  "net/url"
  "encoding/json"
)

type OwnTrade struct {
  Pair string
  Type string
  Amount float32
  Rate float32
  Timestamp int32
}

// Unpacking bullshit
type TradesResponse struct {
  Success int
  Error string
  Return map[string]OwnTrade
}

func decodeTrades(body []byte) TradesResponse {
  var trades TradesResponse
  json.Unmarshal(body, &trades)
  if trades.Success != 1 {
    fmt.Println(trades.Error)
  }
  return trades
}

func LastTradeMade() OwnTrade {
  params := url.Values{}
  params.Set("count", "1")
  responseBody := ApiRequest("TradeHistory", params)
  trades := decodeTrades(responseBody)
  var firstTrade OwnTrade
  for _, trade := range trades.Return {
    firstTrade = trade
    break
  }
  return firstTrade
}

