package analysis

import(
  "data"
)

const giantNumber float32 = 100000000

var hourlyMetrics = [6]int{6,12,24,48,72,120}
var minuteMetrics = [4]int{5,10,30,60}

var statusesCollection = data.GetCollection("statuses")

type Range struct {
  Min float32
  Max float32
}

type Ranges map[string]Range

type Metrics map[string]float32

type Analysis struct {
  Range      Ranges   // Min and max prices
  Percentile Metrics  // Posn of current price within range
  Slope      Metrics  // Overall price trend per hour
  Volatility Metrics  // Average deviation from slope
}

func Analyze(status Status) Status {
  status.Analysis.Range      = calculateRangeMap(status)
  status.Analysis.Percentile = calculatePercentileMap(status)
  status.Analysis.Volatility = calculateVolatilityMap(status)
  status.Analysis.Slope      = calculateSlopeMap(status)
  return status
}


