package troll

import (
  "fmt"
  "btce"
  "market"
)


func (troll Troll) Sell(status market.MarketPrice) Troll {
  rate := status.Price - 0.5 // Go 50 cents in to ensure trade is snatched up
  usd := troll.Funds.BTC * rate

  usd *= 0.998 // fee

  newTrade := btce.OwnTrade{
    Pair: "btc_usd",
    Type: "sell",
    Amount: troll.Funds.BTC,
    Rate: rate,
    Timestamp: status.Time.Server,
  }

  troll.Funds.BTC = 0
  troll.Funds.USD = usd
  troll.LastTrade = newTrade

  fmt.Println("")
  fmt.Println(fmt.Sprintf("\n%sTroll SELL @ %f. USD bal: %f", CLR_WHITE, status.Price, usd))
  fmt.Println("")
  return troll
}

func (troll Troll) Buy(status market.MarketPrice) Troll {
  rate := status.Price + 0.5 // Go 50 cents in to ensure trade is snatched up
  btc := troll.Funds.USD / rate

  btc *= 0.998 // fee

  newTrade := btce.OwnTrade{
    Pair: "btc_usd",
    Type: "buy",
    Amount: troll.Funds.BTC,
    Rate: rate,
    Timestamp: status.Time.Server,
  }

  troll.Funds.BTC = btc
  troll.Funds.USD = 0
  troll.LastTrade = newTrade

  fmt.Println("")
  fmt.Println(fmt.Sprintf("\n%sTroll BUY @ %f. BTC bal: %f", CLR_WHITE, status.Price, btc))
  fmt.Println("")
  return troll
}