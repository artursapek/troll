package main

import (
  "btce"
  "fmt"
  "troll/data"
  "labix.org/v2/mgo/bson"
)

func main () {
  c := data.GetCollection("trades")
  trades := btce.GetTrades()
  for i := 0; i < 150; i ++ {
    trade := trades[i]
    p    := trade.Price
    id   := trade.Tid
    fmt.Println(fmt.Sprintf("%d,%f", id, p))
    change, err := c.Upsert(bson.M{"tid": id}, &trade)
    if err != nil {
      panic(err)
    }
    print(change)
  }

  count, _ := c.Count()
  print(count)
}
