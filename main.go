package main

import (
  "btce"
  "troll/memory"
  "fmt"
)

func main () {
  trades := btce.GetTrades()
  for i := 0; i < 150; i ++ {
    p := trades[i].Price
    id := trades[i].Tid
    fmt.Println(fmt.Sprintf("%d,%f", id, p))
  }
  memory.Append("prices",  "test", "a")

}
