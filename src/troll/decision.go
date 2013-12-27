package troll

import (
  "fmt"
  "analysis"
)

// Troll is great at decision making
func (troll Troll) Decide(status analysis.MarketStatus) Troll {
  //fmt.Printf("%d,%f,%f,%f\n", status.ServerTime, status.Price, status.Analysis.EMA["8"], status.Analysis.EMA["18"])
  //return troll
  if troll.Holding() {
    return troll.DecideWhenHolding(status)
  } else {
    return troll.DecideWhenWaiting(status)
  }
}

func (troll Troll) DecideWhenHolding(status analysis.MarketStatus) Troll {
  //prev, _ := analysis.PreviousStatus(status)

    // Should always be the case, by logic

  if status.Analysis.EMA["8"] < status.Analysis.EMA["18"] {
    // The lines have crossed. Sell
    if status.Price > (troll.LastTrade.Rate * 1.02) {
      return troll.Sell(status)
    }
  }
  fmt.Printf("%sHolding. Rate: $%f\r", CLR_GREY, status.Price)
  return troll
}


func (troll Troll) DecideWhenWaiting(status analysis.MarketStatus) Troll {
  //prev, _ := analysis.PreviousStatus(status)


  if status.Analysis.EMA["8"] > status.Analysis.EMA["18"] {
    // The lines have crossed. Sell
    if status.Price < (troll.LastTrade.Rate * 0.95) {
      return troll.Buy(status)
    }
  }
  fmt.Printf("%sWaiting. Rate: $%f\r", CLR_GREY, status.Price)

  return troll
}

