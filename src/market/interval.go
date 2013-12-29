package market

import (
  "time"
  "data"
  "ichimoku"
)

const INTERVAL_PERIOD = 60 * 60 * 2 // 2 hours in seconds

type MarketInterval struct {
  Time struct { // All local
    Open  int32
    Close int32
  }
  CandleStick CandleStick
  Ichimoku ichimoku.IchimokuLines
}

func LastInterval() MarketInterval {
  var intervals []MarketInterval
  data.Intervals.Find(nil).Sort("-closetime").Limit(1).All(&intervals)
  return intervals[0]
}

func LastIntervalTime() time.Time {
  timestamp := LastInterval().Time.Close
  return time.Unix(int64(timestamp), 0)
}

func timeSinceLastInterval() int32 {
  return int32(time.Now().Unix()) - LastInterval().Time.Close
}

func NewIntervalHasClosed() bool {
  return timeSinceLastInterval() > INTERVAL_PERIOD
}

func RecordInterval() (interval MarketInterval) {
 // now := int32(time.Now().Unix())
  
  return interval
}

