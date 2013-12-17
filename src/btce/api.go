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

