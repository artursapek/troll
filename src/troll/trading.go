package troll

import (
  "fmt"
  "btce"
  "data"
  "market"
)

const CLR_WHITE  = "\x1b[37;1m"
const CLR_GREY   = "\x1b[30;1m"
const CLR_GREEN  = "\x1b[32;1m"
const CLR_YELLOW = "\x1b[33;1m"
const CLR_RED    = "\x1b[31;1m"

func (troll Troll) Sell(interval market.MarketInterval) Troll {
  rate := interval.CandleStick.Close - 0.5 // Go 50 cents in to ensure trade is snatched up
  usd := troll.Funds.BTC * rate

  usd *= 0.998 // fee

  var color string
  if interval.CandleStick.Close < troll.LastTrade.Rate {
    color = CLR_RED
  } else {
    color = CLR_WHITE
  }

  newTrade := btce.OwnTrade{
    Pair: "btc_usd",
    Type: "sell",
    Amount: troll.Funds.BTC,
    Rate: rate,
    Timestamp: interval.Time.Close,
  }

  data.Trades.Insert(&newTrade)

  troll.Funds.BTC = 0
  troll.Funds.USD = usd
  troll.LastTrade = newTrade

  fmt.Println(fmt.Sprintf("%sSELL @ $%f. Bal: $%f B⃦%f", color, interval.CandleStick.Close, usd, troll.Funds.BTC))
  return troll
}

func (troll Troll) Buy(interval market.MarketInterval) Troll {
  rate := interval.CandleStick.Close + 0.5 // Go 50 cents in to ensure trade is snatched up
  btc := troll.Funds.USD / rate

  btc *= 0.998 // fee

  var color string
  if interval.CandleStick.Close > troll.LastTrade.Rate {
    color = CLR_RED
  } else {
    color = CLR_WHITE
  }

  newTrade := btce.OwnTrade{
    Pair: "btc_usd",
    Type: "buy",
    Amount: troll.Funds.BTC,
    Rate: rate,
    Timestamp: interval.Time.Close,
  }

  data.Trades.Insert(&newTrade)

  troll.Funds.BTC = btc
  troll.Funds.USD = 0
  troll.LastTrade = newTrade


  fmt.Println(fmt.Sprintf("%sBUY @ $%f. Bal: $%f B⃦%f", color, interval.CandleStick.Close, troll.Funds.USD, btc))
  return troll
}
