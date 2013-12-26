package analysis

import (
  "strconv"
)

func calculateRange(statuses []MarketStatus) Range {
  min := giantNumber
  max := -giantNumber

  for i := 0; i < len(statuses); i ++ {
    price := statuses[i].Price
    if price > max {
      max = price
    } else if price < min {
      min = price
    }
  }
  return Range{ Min: min, Max: max }
}

func calculateRangeMap(status MarketStatus, statuses []MarketStatus) Ranges {
  r := make(Ranges)

  for i := 0; i < 5; i ++ {
    hrs := hourlyMetrics[i]
    hrsString := strconv.Itoa(hrs)
    r[hrsString] = calculateRange(filterPastNHours(statuses, int32(hrs), status.ServerTime))
  }
  return r
}


