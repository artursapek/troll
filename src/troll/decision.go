package troll

import (
  "market"
)

// Troll is great at decision making
func (troll Troll) Decide(status market.MarketInterval) Troll {
  //fmt.Printf("%d,%f,%f,%f\n", status.ServerTime, status.Price, status.Analysis.EMA["8"], status.Analysis.EMA["18"])
  //return troll
  if troll.Holding() {
    return troll.DecideWhenHolding(status)
  } else {
    return troll.DecideWhenWaiting(status)
  }
}

func (troll Troll) DecideWhenHolding(status market.MarketInterval) Troll {
  return troll
}


func (troll Troll) DecideWhenWaiting(status market.MarketInterval) Troll {
  return troll
}

