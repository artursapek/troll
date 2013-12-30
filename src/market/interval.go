package market

import (
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
  Ichimoku ichimoku.IchimokuLines
}

func LastInterval() MarketInterval {
  var intervals []MarketInterval
  data.Intervals.Find(nil).Sort("-closetime").Limit(1).All(&intervals)
  if len(intervals) == 0 {
    return MarketInterval{}
  } else {
    return intervals[0]
  }
}

func roundUpToNearest2Hour(timestamp int64) int64 {
  t := time.Unix(timestamp, 0)
  tRounded := t.Round(INTERVAL_PERIOD * time.Second).Unix()

  if t.Unix() > tRounded {
    // That means we rounded down, so we want to add another
    // INTERVAL_PERIOD to get a full segment as long as INTERVAL_PERIOD.
    tRounded += INTERVAL_PERIOD
  }

  return tRounded
}

func LastIntervalCloseTime() int64 {
  lastInterval := LastInterval()
  var timestamp int64
  if (lastInterval == MarketInterval{}) {
    // If we have no intervals, just use the first price's time
    timestamp = getFirstPrice().Time.Local
  } else {
    // If it does exist, return its close time like we expected
    timestamp = lastInterval.Time.Close
  }
  return timestamp
}

func TimeSinceLastInterval(now int64) int64 {
  return now - LastIntervalCloseTime()
}

func NewIntervalHasClosed(now int64) bool {
  return TimeSinceLastInterval(now) > INTERVAL_PERIOD
}


func RecordInterval() (interval MarketInterval) {
  // now := int64(time.Now().Unix())
  
  return interval
}

