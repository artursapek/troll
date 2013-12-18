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
  body, readErr := ioutil.ReadAll(res.Body)
  if readErr != nil {
    panic(err)
  }
  return body
}

type Trade struct {
  Date   int32
  Price  float32
  Amount float32
  Tid    int32
}

type Trades []Trade

func decodeTrades(input []byte) Trades {
  var trades Trades
  err := json.Unmarshal(input, &trades)
  if err != nil {
    panic(err)
  }
  return trades
}

func GetTrades() Trades {
  response := PublicApiRequest("trades")
  return decodeTrades(response)
}

type Ticker struct {
  High float32
  Low float32
  Buy float32
  Sell float32
  Server_time int32
}


type TickerUnpack struct {
  Ticker Ticker
}

func decodeTicker(input []byte) Ticker {
  var ticker TickerUnpack
  err := json.Unmarshal(input, &ticker)
  if err != nil {
    panic(err)
  }
  return ticker.Ticker
}

func GetTicker() Ticker {
  response := PublicApiRequest("ticker")
  return decodeTicker(response)
}

