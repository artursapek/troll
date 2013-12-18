package price

import (
  "btce"
  "data"
  "labix.org/v2/mgo/bson"
  "time"
  "fmt"
)

func perform() {
  fmt.Println("Running...")
  c := data.GetCollection("trades")
  trades := btce.GetTrades()
  for i := 0; i < 150; i ++ {
    trade := trades[i]
    tid   := trade.Tid
    change, err := c.Upsert(bson.M{"tid": tid}, &trade)
    if err != nil {
      panic(err)
    }
    print(change)
  }
}

func Run() {
  running := true
  for running {
    perform()
    // Sleep for a minute
    time.Sleep(60 * 1000 * time.Millisecond)
  }
}
