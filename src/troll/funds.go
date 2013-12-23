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


