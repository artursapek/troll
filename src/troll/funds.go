package troll

import (
  "btce"
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

