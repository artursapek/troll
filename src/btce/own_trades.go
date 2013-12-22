package btce

import (
  "net/url"
  "encoding/json"
)

type Trade struct {
  Pair string
  Type string
  Amount float32
  Rate float32
  Order_id int32
  Timestamp int32
}

type TradesResponse struct {
  Success int32
  Return map[string]Trade
}

func DecodeTrades(body []byte) TradesResponse {
  var trades TradesResponse
  json.Unmarshal(body, &trades)
  return trades
}

func LastTrade() Trade {
  params := url.Values{}
  params.Set("count", "1")
  response := ApiRequest("TradeHistory", params)
  trades := DecodeTrades(response)
  var firstTrade Trade
  for _, trade := range trades.Return {
    firstTrade = trade
    break
  }
  return firstTrade
}

