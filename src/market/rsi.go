package market

import (
  "data"
)

const RSIPeriod int = 14

func (interval *MarketInterval) CalculateRSI() {
  var intervals MarketIntervals
  data.Intervals.Find(nil).Limit(RSIPeriod).Sort("-time.close").All(&intervals)

  if len(intervals) < RSIPeriod {
    interval.RSI = 0
    return
  }

  var currentClose, change, gains, losses float32

  lastClose := interval.CandleStick.Close

  for _, intv := range intervals {
    currentClose = intv.CandleStick.Close
    change = lastClose - currentClose

    if (change > 0) {
      gains += change
    } else if (change < 0) {
      losses -= change
    }

    lastClose = currentClose
  }

  interval.RSI = 100 - (100 / (1 + (gains / losses)))
}


