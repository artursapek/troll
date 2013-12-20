package btce

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

const publicApiUrl = "https://btc-e.com/api/2/btc_usd/%s"

func PublicApiRequest(action string) []byte {
  requestUrl := fmt.Sprintf(publicApiUrl, action)
  res, err := http.Get(requestUrl)
  if err != nil {
    panic(err)
  }
  body, _ := ioutil.ReadAll(res.Body)
  return body
}

type PublicTrade struct {
  Date   int32
  Price  float32
  Amount float32
  Tid    int32
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

type Ticker struct {
  High float32
  Low float32
  Buy float32
  Sell float32
  Price float32
  Server_time int32
}


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

