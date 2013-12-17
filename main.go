package main

import (
  "btce"
  "fmt"
)

func main () {
  //res := btce.ApiRequest("TradeHistory", values)
  //body, _ := ioutil.ReadAll(res.Body)
  //print(string(body))
  trades := btce.GetTrades()
  var lastPrice float32 = 0
  for i := 0; i < 150; i ++ {
    p := trades[i].Price
    fmt.Println(p - lastPrice)
    lastPrice = p
  }
}
