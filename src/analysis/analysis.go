package analysis

import(
  "fmt"
  "data"
  "math"
  "labix.org/v2/mgo/bson"
)

var coll = data.GetCollection("statuses")
var hoursToCollect = [5]int{12,24,48,72,120}

type Range struct {
  Min float32
  Max float32
}

type Ranges map[int]Range

type Metrics map[int]float32

type Analysis struct {
  Range      Ranges   // Min and max prices
  Percentile Metrics  // Posn of current price within range
  Slope      Metrics  // Overall price trend per hour
  Volatility Metrics  // Average deviation from slope
}

func Analyze(status Status) (analysis Analysis) {
  analysis.Range      = calculateRangeMap(status)
  analysis.Percentile = calculatePercentileMap(status)
  analysis.Volatility = calculateVolatilityMap(status)
  analysis.Slope      = calculateSlopeMap(status)
  return analysis
}

func statusesFromPast7Days (status Status) (statuses []Status) {
  // 7 days in seconds
  start := status.Time - (7 * 24 * 60 * 60)
  // Find all statuses between the two dates, and unpack them into statuses variable
  query := bson.M{ "time": bson.M{ "$gte": start, "$lt": status.Time }}
  err := coll.Find(query).Sort("-time").All(&statuses)
  if err != nil {
    fmt.Println(err) // Log it, and return empty
  }
  fmt.Println(fmt.Sprintf("Found %d past statuses", len(statuses)))
  fmt.Println(statuses)
  return statuses
}

func pastNHours (statuses []Status, n, now int32) (results []Status) {
  start := now - (n * 60 * 60)
  for i := 0; i < len(statuses); i ++ {
    status := statuses[i]
    if status.Time > start {
      results = append(results, status)
    }
  }
  return results
}

// Range

func calculateRange(statuses []Status) (r Range) {
  min := float32(math.Inf(1))
  max := float32(math.Inf(-1))

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

func calculateRangeMap(status Status) (r Ranges) {
  statuses := statusesFromPast7Days(status)
  for i := 0; i < 5; i ++ {
    hrs := hoursToCollect[i]
    r[hrs] = calculateRange(pastNHours(statuses, int32(hrs), status.Time))
  }
  return r
}

// Percentile

func calculatePercentileMap(status Status) (r Metrics) {
  statuses := statusesFromPast7Days(status)
  return r
}

// Volatility

func calculateVolatilityMap(status Status) (r Metrics) {
  statuses := statusesFromPast7Days(status)
  return r
}

// Slope

func calculateSlopeMap(status Status) (r Metrics) {
  statuses := statusesFromPast7Days(status)
  return r
}
