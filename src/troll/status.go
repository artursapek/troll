package troll

import (
  "fmt"
  "data"
  "time"
  "market"
)

func LastUpdate() {
  now := time.Now().Unix()
  var unpack []market.MarketPrice
  data.Prices.Find(nil).Sort("-time.local").Limit(1).All(&unpack)
  fmt.Println(now - unpack[0].Time.Local)
}
