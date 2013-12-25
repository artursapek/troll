package troll

import (
  "fmt"
  "analysis"
)

// Get more aggressive about buying after a while
const NervousnessThreshold = 14

// Troll is great at decision making

func (troll Troll) Decide(status analysis.MarketStatus) Troll {
  fmt.Printf(".")
  if troll.Holding() {
    return troll.DecideWhenHolding(status)
  } else {
    return troll.DecideWhenWaiting(status)
  }
}

func (troll Troll) DecideWhenHolding(status analysis.MarketStatus) Troll {
  thresholdProfit := troll.SellThreshold(status)
  potentialProfit := troll.PotentialProfit(status)

  if status.Analysis.Slope.Accelerating() {
    //fmt.Println(//fmt.Sprintf("Too severe: %f     ", status.Analysis.Slope["5"]))
    return troll
  }

  if status.Analysis.Percentile["all"] > 0.995 {
    // Hold on if by some miracle the value is at an all-time high
    //fmt.Println("Too valuable")
    return troll
  }

  if potentialProfit >= thresholdProfit {
    if status.Analysis.Percentile["6"] < 0.2 {

      if !status.Analysis.Slope.Accelerating() {
        //fmt.Println("Sell b/c good margin, low percentile, has settled")
        return troll.Sell(status)
      } else {
        //fmt.Println("Hasnt settled")
      }
    } else {
      //fmt.Println("Percentile too high")
    }
  } else {
    //fmt.Println(//fmt.Sprintf("Bad margin: %f %f (%f,%f)", potentialMargin, thresholdMargin, status.Price, troll.LastTrade.Rate))
  }
  return troll
}


func (troll Troll) DecideWhenWaiting(status analysis.MarketStatus) Troll {
  thresholdProfit := troll.BuyThreshold(status)
  potentialProfit := troll.PotentialProfit(status)

  if status.Analysis.Slope.Accelerating() {
    // Never trade while the trend is severe.
    // Wait until it settles somewhat.
    //fmt.Println(//fmt.Sprintf("Too severe: %f", status.Analysis.Slope["5"]))
    return troll
  }
  ////fmt.Println(threshold)

  if status.Analysis.Percentile["all"] > 0.995 && troll.WaitingForTooLong(status) {
    if (potentialProfit / troll.LastTrade.Rate) > -0.05 {
      ////fmt.Println("RECORD HIGH AND LOW LOSSES. BUYING OMG")
      // BUY
      //fmt.Println("Record high, buying")
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
        return troll.Buy(status)
      } else {
        //fmt.Println(//fmt.Sprintf("Accelerating; waiting", potentialMargin, thresholdMargin))
      }
    } else {
      //fmt.Println(//fmt.Sprintf("Bad margin: %f %f (%f,%f)", potentialMargin, thresholdMargin, status.Price, troll.LastTrade.Rate))
    }
  }

  return troll
}

func (troll Troll) TimeWaiting(status analysis.MarketStatus) int32 {
  return status.ServerTime - troll.LastTrade.Timestamp
}

func (troll Troll) WaitingForTooLong(status analysis.MarketStatus) bool {
  return (troll.TimeWaiting(status) / 60 / 60) > NervousnessThreshold
}

