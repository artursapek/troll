package market

import (
  "data"
  "mathutils"
  "labix.org/v2/mgo/bson"
)

// Calculation periods (inclusive to current interval)
const tenkanPeriod int = 7  // Traditionally 9
const kijunPeriod  int = 11 // Traditionally 26

type Indicators struct {
  TenkanSen   float32
  KijunSen    float32
  SenkouSpanA float32
  SenkouSpanB float32
  ChikouSpan  float32 // Close price from kijunPeriod ago
}

func CalculateIndicators(interval MarketInterval) (indicators Indicators) {
  // Hey cuz, get the intervals we need cuz
  intervals := GetIchimokuIntervalsUntil(interval)

  indicators.TenkanSen   = MidPointOfHighLow(intervals, tenkanPeriod)
  indicators.KijunSen    = MidPointOfHighLow(intervals, kijunPeriod)
  indicators.SenkouSpanA = (indicators.KijunSen + indicators.TenkanSen) / 2
  indicators.SenkouSpanB = MidPointOfHighLow(intervals, kijunPeriod * 2)
  indicators.ChikouSpan  = LaggingChikou(intervals)
  return indicators
}

func UnpackIchimoku(i Indicators) (float32, float32, float32, float32, float32) {
  return i.TenkanSen, i.KijunSen, i.SenkouSpanA, i.SenkouSpanB, i.ChikouSpan
}

func MidPointOfHighLow(intervals MarketIntervals, n int) float32 {
  var min, max float32

  for i, interval := range intervals {
    if i == n {
      break // Don't go further than n intervals =^_^=
    } else if i == 0 {
      min = interval.CandleStick.Low
      max = interval.CandleStick.High
    } else {
      min = mathutils.Min(min, interval.CandleStick.Low)
      max = mathutils.Max(max, interval.CandleStick.High)
    }
  }

  return (min + max) / 2
}

func LaggingChikou(intervals MarketIntervals) (chikou float32) {
  // The price we want is kijunPeriod periods ago,
  // so the index of that interval is (kijunPeriod - 1)
  indexAt := kijunPeriod - 1
  for i, interval := range intervals {
    if i == indexAt {
      chikou = interval.CandleStick.Close
      break
    }
  }
  return chikou
}

func GetIchimokuIntervalsUntil (interval MarketInterval) MarketIntervals {
  // I feel like there is a better way of doing this

  var intervals MarketIntervals

  // We need the past (kijunPeriod * 2) intervals, including the current one,
  // to perform the Ichimoku analysis. Since we get the current one passed
  // in before it's persisted to Mongo, we need to query for (kijunPeriod * 2) - 1
  n := (kijunPeriod * 2) - 1

  // Run the query and unpack into throwaway slice
  var throwaway []MarketInterval
  query := bson.M{ "time.close": bson.M{ "$lt": interval.Time.Close }}
  data.Intervals.Find(query).Sort("-time.close").Limit(n).All(&throwaway)

  // First element is the last interval we passed in
  // since we're decrementing by time.close
  intervals = append(intervals, interval)

  // Then append each of the persisted intervals
  // (still in decrementing order)
  for _, in := range throwaway {
    intervals = append(intervals, in)
  }

  return intervals
}

