package analysis

import (
  "strconv"
)

func calculatePercentileMap(status MarketStatus) Metrics {
  metrics := make(Metrics)
  for i := 0; i < 5; i ++ {
    hrs := hourlyMetrics[i]
    hrsString := strconv.Itoa(hrs)
    r := status.Analysis.Range[hrsString]
    d := r.Max - r.Min
    perc := (status.Price - r.Min) / d
    if perc < 0 {
      perc *= -1
    }
    metrics[hrsString] = perc
  }
  return metrics
}


