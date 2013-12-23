package troll

import (
  "fmt"
  "analysis"
)

// Troll is great at decision making

func (troll Troll) Decide(status analysis.MarketStatus) {
  fmt.Println(status.Price)
  fmt.Println(troll.LastTrade.Rate > status.Price)
}

