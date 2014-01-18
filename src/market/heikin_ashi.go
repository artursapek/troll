package market

import (
  "mathutils"
)

func (interval *MarketInterval) CalculateHeikinAshi() {
  prev := interval.Prev()
  candle := interval.CandleStick
  ha := CandleStick{}

  ha.Close = (candle.Open + candle.Close + candle.High + candle.Low) / 4

  if (prev == MarketInterval{}) {
    // Base case
    ha.Open = (candle.Open + candle.Close) / 2
    ha.High = candle.High
    ha.Low  = candle.Low
  } else {
    pha := prev.HeikinAshi
    ha.Open = (pha.Open + pha.Close) / 2
    ha.High = mathutils.Max(candle.High, pha.Open, pha.Close)
    ha.Low  = mathutils.Min(candle.Low, pha.Open, pha.Close)
  }

  interval.HeikinAshi = ha
}
