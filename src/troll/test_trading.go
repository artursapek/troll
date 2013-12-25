package troll

import (
  "fmt"
  "btce"
  "analysis"
)

func (troll Troll) Sell(status analysis.MarketStatus) Troll {
  rate := status.Price - 0.5 // Go 50 cents in to ensure trade is snatched up
  usd := troll.Funds.BTC * rate

  newTrade := btce.OwnTrade{
    Pair: "btc_usd",
    Type: "sell",
    Amount: troll.Funds.BTC,
    Rate: rate,
    Timestamp: status.ServerTime,
  }

  troll.Funds.BTC = 0
  troll.Funds.USD = usd
  troll.LastTrade = newTrade

  fmt.Println(fmt.Sprintf("\nTroll sold at %f. USD bal: %f        ", status.Price, usd))
  return troll
}

func (troll Troll) Buy(status analysis.MarketStatus) Troll {
  rate := status.Price + 0.5 // Go 50 cents in to ensure trade is snatched up
  btc := troll.Funds.USD / rate

  newTrade := btce.OwnTrade{
    Pair: "btc_usd",
    Type: "buy",
    Amount: troll.Funds.BTC,
    Rate: rate,
    Timestamp: status.ServerTime,
  }

  troll.Funds.BTC = btc
  troll.Funds.USD = 0
  troll.LastTrade = newTrade

  fmt.Println(fmt.Sprintf("\nTroll bought at %f. BTC bal: %f        ", status.Price, btc))
  return troll
}
