package troll

import (
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
  potentialMargin := troll.PotentialMargin(status)

  threshRange := status.Analysis.Range["12"]
  // Relative threshold based on recent trade history
  thresholdMargin := ((threshRange.Max - threshRange.Min) / troll.LastTrade.Rate) / float32(2)

  if status.SlopeIsSevere() {
    //fmt.Println(//fmt.Sprintf("Too severe: %f     ", status.Analysis.Slope["5"]))
    return troll
  }

  if status.Analysis.Percentile["all"] > 0.995 {
    // Never trade while the trend is severe.
    // Wait until it settles somewhat.
    //fmt.Println("Too valuable")
    return troll
  }

  if potentialMargin > thresholdMargin {
    if status.Analysis.Percentile["6"] < 0.2 {

      if !status.Analysis.Slope.IsAccelerating() {
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
  potentialMargin := troll.PotentialMargin(status)

  threshRange := status.Analysis.Range["12"]
  thresholdMargin := ((threshRange.Max - threshRange.Min) / troll.LastTrade.Rate) / float32(2.5)

  ////fmt.Println(potentialProfit)

  if status.SlopeIsSevere() {
    // Never trade while the trend is severe.
    // Wait until it settles somewhat.
    //fmt.Println(//fmt.Sprintf("Too severe: %f", status.Analysis.Slope["5"]))
    return troll
  }
  ////fmt.Println(threshold)

  if status.Analysis.Percentile["all"] > 0.995 && troll.WaitingForTooLong(status) {

    if potentialMargin > -0.05 {
      ////fmt.Println("RECORD HIGH AND LOW LOSSES. BUYING OMG")
      // BUY
      //fmt.Println("Record high, buying")
      return troll.Buy(status)
    }
  } else {

    if troll.WaitingForTooLong(status) {
      // Get hasty after a while. Want to be holding.
      thresholdMargin *= float32(0.5)
    }

    if potentialMargin > thresholdMargin {

      if !status.Analysis.Slope.IsAccelerating() {
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

func (troll Troll) PotentialMargin(status analysis.MarketStatus) float32 {
  var diff float32
  lastRate := troll.LastTrade.Rate
  if troll.Holding() {
    diff = status.Price - lastRate
  } else {
    diff = lastRate - status.Price
  }
  return diff / lastRate
}

func (troll Troll) TimeWaiting(status analysis.MarketStatus) int32 {
  return status.ServerTime - troll.LastTrade.Timestamp
}

func (troll Troll) WaitingForTooLong(status analysis.MarketStatus) bool {
  return (troll.TimeWaiting(status) / 60 / 60) > NervousnessThreshold
}

