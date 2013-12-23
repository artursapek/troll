package analysis

import(
  "data"
)

const giantNumber float32 = 100000000

var hourlyMetrics = [5]int{6,12,24,48,72} // Long-term
var minuteMetrics = [4]int{5,10,30,60}    // Short-term

var statusesCollection = data.GetCollection("test_prices")

type Range struct {
  Min, Max float32
}

type Ranges map[string]Range

type Metrics map[string]float32

type Analysis struct {
  Range      Ranges  // Min and max prices; long term
  Percentile Metrics // Posn of current price within range; long term
  Slope      Metrics // Overall price trend per hour; short term
  Volatility Metrics // Average deviation from slope; long term
}

func Analyze(status MarketStatus) MarketStatus {
  status.Analysis.Range      = calculateRangeMap(status)
  status.Analysis.Percentile = calculatePercentileMap(status)
  status.Analysis.Volatility = calculateVolatilityMap(status)
  status.Analysis.Slope      = calculateSlopeMap(status)
  return status
}

