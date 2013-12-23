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



// "Holding" in this app means holding BTC.
// "Waiting" means waiting to buy BTC back.
// The implicit names are just for convenience,
// and to emphasize the bot's preference
// for having BTC as opposed to USD.

func (status FundsStatus) Holding() bool {
  return status.BTC > status.USD
}

func (status FundsStatus) Waiting() bool {
  return status.BTC < status.USD
}
