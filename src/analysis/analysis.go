package analysis

import(
  "fmt"
  "data"
  "strconv"
  "labix.org/v2/mgo/bson"
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

func statusesFromPast7Days (status Status) (statuses []Status) {
  // 7 days in seconds
  start := status.ServerTime - (7 * 24 * 60 * 60)
  // Find all statuses between the two dates, and unpack them into statuses variable
  query := bson.M{ "servertime": bson.M{ "$gte": start, "$lt": status.ServerTime }}
  err := statusesCollection.Find(query).Sort("-time").All(&statuses)
  if err != nil {
    fmt.Println(err) // Log it, and return empty
  }
  fmt.Println(fmt.Sprintf("Found %d past statuses", len(statuses)))
  return statuses
}

func pastNHours (statuses []Status, n, now int32) (results []Status) {
  start := now - (n * 60 * 60)
  for i := 0; i < len(statuses); i ++ {
    status := statuses[i]
    if status.ServerTime > start {
      results = append(results, status)
    }
  }
  return results
}

func pastNMinutes (statuses []Status, n, now int32) (results []Status) {
  start := now - (n * 60)
  for i := 0; i < len(statuses); i ++ {
    status := statuses[i]
    if status.ServerTime > start {
      results = append(results, status)
    }
  }
  return results
}


// Range

func calculateRange(statuses []Status) Range {
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

func calculateRangeMap(status Status) Ranges {
  r := make(Ranges)
  statuses := statusesFromPast7Days(status)
  for i := 0; i < 5; i ++ {
    hrs := hourlyMetrics[i]
    hrsString := strconv.Itoa(hrs)
    r[hrsString] = calculateRange(pastNHours(statuses, int32(hrs), status.ServerTime))
  }
  return r
}

// Percentile

func calculatePercentileMap(status Status) Metrics {
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

// Volatility

func calculateVolatility(statuses []Status) (avgDev float32) {
  amt := len(statuses)
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

func calculateVolatilityMap(status Status) Metrics {
  metrics := make(Metrics)
  statuses := statusesFromPast7Days(status)
  if len(statuses) == 0 {
    return metrics
  }
  for i := 0; i < 5; i ++ {
    hrs := hourlyMetrics[i]
    hrsString := strconv.Itoa(hrs)
    metrics[hrsString] = calculateVolatility(pastNHours(statuses, int32(hrs), status.ServerTime))
  }
  return metrics
}

func calculateSlope(statuses []Status) (slope float32) {
  amt := len(statuses)
  amtFloat := float32(amt)
  first := statuses[0]
  last := statuses[amt - 1]
  fmt.Println(first.Price)
  fmt.Println(last.Price)
  return (last.Price - first.Price) / amtFloat
}

func calculateSlopeMap(status Status) Metrics {
  metrics := make(Metrics)
  statuses := statusesFromPast7Days(status)
  if len(statuses) == 0 {
    return metrics
  }
  for i := 0; i < 4; i ++ {
    mins := minuteMetrics[i]
    minsString := strconv.Itoa(mins)
    metrics[minsString] = calculateSlope(pastNMinutes(statuses, int32(mins), status.ServerTime))
  }
  return metrics
}



// Slope

/*
func calculateSlopeMap(status Status) (r Metrics) {
  statuses := statusesFromPast7Days(status)
  return r
}
*/
