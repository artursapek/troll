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
  RSI         float32
  TR          float32
  ATR         float32
  EMA10       float32
  EMA21       float32
  CandleStick CandleStick
  HeikinAshi  CandleStick
  Ichimoku    Indicators
}

type MarketIntervals []MarketInterval

// Creator

func RecordIntervalSucceeding(prevInterval MarketInterval) (interval MarketInterval) {
  openTime  := prevInterval.Time.Close
  closeTime := openTime + INTERVAL_PERIOD

  lastClosePrice := prevInterval.CandleStick.Close

  prices := getPricesBetween(openTime, closeTime)
  return RecordIntervalFromPrices(prices, openTime, lastClosePrice)
}

func RecordIntervalFromPrices(prices []MarketPrice, openTime int64, lastClosePrice float32) (interval MarketInterval) {
  interval.Time.Open = openTime
  interval.Time.Close = openTime + INTERVAL_PERIOD

  interval.CalculateCandleStick(prices, lastClosePrice)

  interval.Analyze()

  // Persist to db
  data.Intervals.Insert(&interval)

  fmt.Printf(".")

  return interval
}

func (interval *MarketInterval) Analyze() {
  interval.CalculateHeikinAshi()
  interval.CalculateRSI()
  interval.CalculateATR()
  interval.CalculateEMAs()
  interval.CalculateParabolicSAR()
  interval.CalculateIchimokuIndicators()
}

func PersistUpdatedInterval(interval MarketInterval) {
  query  := bson.M{ "time.close": interval.Time.Close }
  data.Intervals.Update(query, interval)
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

func CountOfIntervalsBefore(interval MarketInterval) int {
  count, err := data.Intervals.Find(nil).Sort("-time.close").Count()
  if err != nil {
    panic(err)
  }
  return count
}

func LastInterval() MarketInterval {
  intervalUnpack := PastNIntervals(1)
  return intervalUnpack[0]
}

func (interval MarketInterval) Prev() MarketInterval {
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

