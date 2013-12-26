package troll

import (
  "fmt"
  "analysis"
)

// Get more aggressive about buying after a while
const NervousnessThreshold = 14

// Troll is great at decision making

func (troll Troll) Decide(status analysis.MarketStatus) Troll {
  if troll.Holding() {
    return troll.DecideWhenHolding(status)
  } else {
    return troll.DecideWhenWaiting(status)
  }
}

func (troll Troll) DecideWhenHolding(status analysis.MarketStatus) Troll {
  thresholdProfit := troll.SellThreshold(status)
  potentialProfit := troll.PotentialProfit(status)

  absVola := status.Analysis.Volatility["6"]
  if absVola < 0 {
    absVola *= -1
  }

  // Don't sell when there's no volatility,
  // meaning things are flat and stable.
  // We need whales to buy back cheaper ASAP
  if absVola < 1.0 {
    fmt.Printf("v")
    return troll
  }

  if status.Analysis.Slope.Accelerating() {
    //fmt.Println(//fmt.Sprintf("Too severe: %f     ", status.Analysis.Slope["5"]))
    fmt.Printf("a")
    return troll
  }

  if status.Analysis.Percentile["all"] > 0.995 {
    // Hold on if by some miracle the value is at an all-time high
    //fmt.Println("Too valuable")
    fmt.Printf("!")
    return troll
  }

  perc := status.Analysis.Percentile

  if potentialProfit >= thresholdProfit {
    if perc["6"] < 0.2 && (perc["6"] > perc["12"]) {

      fmt.Printf("p")
      return troll.Sell(status)
    } else {
      //fmt.Println("Percentile too high")
    }
  } else {
    //fmt.Println(//fmt.Sprintf("Bad margin: %f %f (%f,%f)", potentialMargin, thresholdMargin, status.Price, troll.LastTrade.Rate))
  }
  fmt.Printf(".")
  return troll
}


func (troll Troll) DecideWhenWaiting(status analysis.MarketStatus) Troll {
  thresholdProfit := troll.BuyThreshold(status)
  potentialProfit := troll.PotentialProfit(status)

  if status.Analysis.Slope.Accelerating() {
    // Never trade while the trend is severe.
    // Wait until it settles somewhat.
    //fmt.Println(//fmt.Sprintf("Too severe: %f", status.Analysis.Slope["5"]))
    fmt.Printf("A")
    return troll
  }

  ////fmt.Println(threshold)

  if status.Analysis.Percentile["all"] > 0.995 && troll.WaitingForTooLong(status) {
    if (potentialProfit / troll.LastTrade.Rate) > -0.05 {
      ////fmt.Println("RECORD HIGH AND LOW LOSSES. BUYING OMG")
      // BUY
      //fmt.Println("Record high, buying")
      fmt.Printf("R")
      return troll.Buy(status)
    }
  } else {

    if troll.WaitingForTooLong(status) {
      // Get hasty after a while. Want to be holding.
      thresholdProfit /= 2
    }

    if potentialProfit >= thresholdProfit {

      if !status.Analysis.Slope.Accelerating() {
        //fmt.Println("Buying b/c of good margin and it has settled")
        fmt.Printf("P")
        return troll.Buy(status)
      } else {
        //fmt.Println(//fmt.Sprintf("Accelerating; waiting", potentialMargin, thresholdMargin))
      }
    } else {
      //fmt.Println(//fmt.Sprintf("Bad margin: %f %f (%f,%f)", potentialMargin, thresholdMargin, status.Price, troll.LastTrade.Rate))
    }
  }

  fmt.Printf("-")
  return troll
}

func (troll Troll) TimeWaiting(status analysis.MarketStatus) int32 {
  return status.ServerTime - troll.LastTrade.Timestamp
}

func (troll Troll) WaitingForTooLong(status analysis.MarketStatus) bool {
  return (troll.TimeWaiting(status) / 60 / 60) > NervousnessThreshold
}

