package market

import (
  "fmt"
  "mathutils"
  "data"
  "labix.org/v2/mgo/bson"
)

const INTERVAL_PERIOD = 60 * 60 * 2 // 2 hours in seconds

type MarketInterval struct {
  Time struct { // All local
    Open  int64
    Close int64
  }
  SAR         ParabolicSAR
  CandleStick CandleStick
  Ichimoku    Indicators
  Position    string // "long" or "short"
}

type MarketIntervals []MarketInterval

// Creator

func RecordInterval(openTime int64) (interval MarketInterval) {
  closeTime := openTime + INTERVAL_PERIOD
  prices := getPricesBetween(openTime, closeTime)

  interval.Time.Open = openTime
  interval.Time.Close = closeTime
  interval.CandleStick = createCandleStick(prices)

  // Persist to db
  data.Intervals.Insert(&interval)

  interval = AnalyzeInterval(interval)

  return interval
}

func AnalyzeInterval(interval MarketInterval) MarketInterval {
   // Calculating the SAR
  prev := PrevInterval(interval)
  prevPrev := PrevInterval(prev)
  // It comes out sorted by time decrementing
  interval.SAR = CalculateParabolicSAR(interval, prev, prevPrev)

  // Calculating the Ichimoku indicators for this interval
  interval.Ichimoku = CalculateIndicators(interval)

  return interval
}

func SetPosition(interval MarketInterval, position string) {
  if position != "long" && position != "short" {
    panic("Invalid position")
  } else {
    query  := bson.M{ "time.close": interval.Time.Close }
    update := bson.M{ "$set": bson.M { "position": position }}
    data.Intervals.Update(query, update)
  }
}

// Helpers

func lastIntervalCloseTime() int64 {
  intervals := PastNIntervals(1)
  if len(intervals) == 1 {
    // If it does exist, return its close time like we expected
    return intervals[0].Time.Close
  } else {
    // If we have no intervals, just use the first price's time
    timestamp := getFirstPrice().Time.Local
    return mathutils.RoundUpToNearestInterval(timestamp, INTERVAL_PERIOD)
  }
}

func CheckIfNewIntervalIsDue(currentTime int64) (int64, bool) {
  lastClose := lastIntervalCloseTime()
  return lastClose, (currentTime - lastClose) >= INTERVAL_PERIOD
}



// Getters

func PastNIntervals(n int) (intervals MarketIntervals) {
  data.Intervals.Find(nil).Sort("-time.close").Limit(n).All(&intervals)
  return intervals
}

func PrevInterval(interval MarketInterval) MarketInterval {
  var intervals MarketIntervals
  query := bson.M{ "time.close": bson.M{ "$lt": interval.Time.Close } }
  data.Intervals.Find(query).Sort("-time.close").Limit(1).All(&intervals)

  if len(intervals) == 0 { return MarketInterval{} }
  return intervals[0]
}

func NextInterval(interval MarketInterval) MarketInterval {
  var intervals MarketIntervals
  query := bson.M{ "time.close": bson.M{ "$gt": interval.Time.Close } }
  data.Intervals.Find(query).Sort("time.close").Limit(1).All(&intervals)

  if len(intervals) == 0 { return MarketInterval{} }
  return intervals[0]
}

