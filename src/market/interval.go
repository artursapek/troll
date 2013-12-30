package market

import (
  "fmt"
  "time"
  "data"
  "ichimoku"
)

const INTERVAL_PERIOD = 60 * 60 * 2 // 2 hours in seconds

type MarketInterval struct {
  Time struct { // All local
    Open  int64
    Close int64
  }
  CandleStick CandleStick
  Ichimoku ichimoku.Indicators
}

func RecordInterval(openTime int64) (interval MarketInterval) {
  closeTime := openTime + INTERVAL_PERIOD
  prices := getPricesBetween(openTime, closeTime)

  fmt.Println(len(prices))

  interval.Time.Open = openTime
  interval.Time.Close = closeTime
  interval.CandleStick = createCandleStick(prices)

  data.Intervals.Insert(&interval)

  return interval
}

// Helpers

func roundUpToNearest2Hour(timestamp int64) int64 {
  t := time.Unix(timestamp, 0)
  tRounded := t.Round(INTERVAL_PERIOD * time.Second).Unix()

  if t.Unix() > tRounded {
    // That means we rounded down, and we wanted to round up
    tRounded += INTERVAL_PERIOD
  }

  return tRounded
}

func lastIntervalCloseTime() int64 {
  intervals := pastNIntervals(1)
  if len(intervals) == 1 {
    // If it does exist, return its close time like we expected
    return intervals[0].Time.Close
  } else {
    // If we have no intervals, just use the first price's time
    timestamp := getFirstPrice().Time.Local
    return roundUpToNearest2Hour(timestamp)
  }
}

func pastNIntervals(n int) (intervals []MarketInterval) {
  data.Intervals.Find(nil).Sort("-time.close").Limit(n).All(&intervals)
  return intervals
}

