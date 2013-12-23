package analysis

import (
  "strconv"
)

// Volatility

func calculateVolatility(statuses []MarketStatus) (avgDev float32) {
  amt := len(statuses)
  if amt == 0 {
    return 0
  }
  amtFloat := float32(amt)
  first := statuses[0]
  last := statuses[amt - 1]
  slope := (last.Price - first.Price) / amtFloat // Average in/decrease per status

  for i := 0; i < amt; i ++ {
    priceOnFlatGrowth := first.Price + (float32(i) * slope)
    dev := statuses[i].Price - priceOnFlatGrowth
    avgDev += dev / amtFloat
  }

  return avgDev
}

func calculateVolatilityMap(status MarketStatus) Metrics {
  metrics := make(Metrics)
  statuses := statusesFromPastNHours (status, 7 * 24)
  if len(statuses) == 0 {
    return metrics
  }
  for i := 0; i < 5; i ++ {
    hrs := hourlyMetrics[i]
    hrsString := strconv.Itoa(hrs)
    metrics[hrsString] = calculateVolatility(filterPastNHours(statuses, int32(hrs), status.ServerTime))
  }
  return metrics
}


