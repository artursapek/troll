package analysis

import (
  "strconv"
)

func calculateSlope(statuses []MarketStatus) (slope float32) {
  amt := len(statuses)
  if amt == 0 {
    return 0.0
  }
  first := statuses[0]      // Most distant
  last := statuses[amt - 1] // Most recent
  return first.Price - last.Price
}

func calculateSlopeMap(status MarketStatus, statuses []MarketStatus) Metrics {
  metrics := make(Metrics)
  if len(statuses) == 0 {
    return metrics
  }
  for i := 0; i < 4; i ++ {
    mins := minuteMetrics[i]
    minsString := strconv.Itoa(mins)
    filteredStatuses := filterPastNMinutes(statuses, int32(mins), status.ServerTime)
    metrics[minsString] = calculateSlope(filteredStatuses)
  }
  return metrics
}

