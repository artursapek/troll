package analysis

import (
  "strconv"
)

func calculateSlope(statuses []Status) (slope float32) {
  if len(statuses) == 0 {
    return 0.0
  }
  amt := len(statuses)
  first := statuses[0]      // Most distant
  last := statuses[amt - 1] // Most recent
  return first.Price - last.Price
}

func calculateSlopeMap(status Status) Metrics {
  metrics := make(Metrics)
  statuses := statusesFromPastNHours (status, 1)
  if len(statuses) == 0 {
    return metrics
  }
  for i := 0; i < 4; i ++ {
    mins := minuteMetrics[i]
    minsString := strconv.Itoa(mins)
    metrics[minsString] = calculateSlope(filterPastNMinutes(statuses, int32(mins), status.ServerTime))
  }
  return metrics
}

