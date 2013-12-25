package troll

import (
  "btce"
  "analysis"
)

type FundsStatus struct {
  BTC float32
  USD float32
}

func GetFundsStatus() (fundsStatus FundsStatus) {

  funds := btce.GetFunds()

  fundsStatus.BTC = funds["btc"]
  fundsStatus.USD = funds["usd"]

  return fundsStatus
}

// "Holding" means holding BTC.
// "Waiting" means waiting to buy BTC back.
// The implicit names are just for convenience,
// and to emphasize troll's preference
// for having BTC as opposed to USD.

func (troll Troll) Holding() bool {
  return troll.Funds.BTC > troll.Funds.USD
}

func (troll Troll) Waiting() bool {
  return troll.Funds.BTC < troll.Funds.USD
}

func (troll Troll) SellThreshold(status analysis.MarketStatus) float32 {
  r := status.Analysis.Range["12"]
  d := r.Max - r.Min
  return d / 3
}

func (troll Troll) BuyThreshold(status analysis.MarketStatus) float32 {
  r := status.Analysis.Range["12"]
  d := r.Max - r.Min
  // Lower threshold for buying back in
  return d / 3.5
}

func (troll Troll) PotentialProfit(status analysis.MarketStatus) float32 {
  var profit float32
  if troll.Holding() {
    profit = status.Price - troll.LastTrade.Rate
  } else {
    profit = troll.LastTrade.Rate - status.Price
  }
  return profit
}

func (troll Troll) Excited(status analysis.MarketStatus) bool {
  if status.Analysis.Slope.Accelerating() {
    return true
  }
  if troll.Holding() {
    // The potential profit is within 90% of the thresdhold price, or
    // the slope of the market right now is steep.
    return (troll.PotentialProfit(status) / troll.SellThreshold(status)) > 0.90
  } else {
    return (troll.PotentialProfit(status) / troll.BuyThreshold(status)) > 0.90
  }
}

