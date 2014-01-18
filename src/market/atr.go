package market

import (
  "mathutils"
)

const ATRPeriod int = 14
const ATRPeriodFloat float32 = 14

func (interval *MarketInterval) CalculateATR() {
  var trMethod1, trMethod2, trMethod3 float32

  prev := interval.Prev()

  if (prev == MarketInterval{}) {
    return
  }

  trMethod1 = interval.CandleStick.High - interval.CandleStick.Low
  trMethod2 = mathutils.Abs(interval.CandleStick.High - prev.CandleStick.Close)
  trMethod3 = mathutils.Abs(interval.CandleStick.Low - prev.CandleStick.Close)

  interval.TR = mathutils.Max(trMethod1, trMethod2, trMethod3)

  countBefore := CountOfIntervalsBefore(*interval)

  if countBefore == ATRPeriod {
    // The first ATR value is on the fifteenth interval:
    // it is the average of the past 14 intervals' TR values.

    // WARNING this call relies on the current interval
    // not being persisted in Mongo yet.
    prevIntervals := PastNIntervals(ATRPeriod)
    var atr float32 = 0
    for _, intvl := range prevIntervals {
      atr += intvl.TR
    }
    atr /= ATRPeriodFloat
    interval.ATR = atr
  } else if countBefore > ATRPeriod {
    // Every ATR after that involves the previous ATR
    interval.ATR = ((prev.ATR * 13) + interval.TR) / ATRPeriodFloat
  }
  // ... else we don't assign ATR yet - we need
  // at least 14 past intervals first.
}
