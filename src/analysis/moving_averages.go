package analysis

const SmooCo float32 = 0.015

func calculateEMA(status MarketStatus, pastStatuses []MarketStatus) float32 {
  prev := pastStatuses[0]

  var prevEMA float32

  if prev.Analysis.EMA == 0 {
    prevEMA = 1
  } else {
    prevEMA = prev.Analysis.EMA
  }

  return SmooCo * prev.Price + (1 - SmooCo) * prevEMA
}

