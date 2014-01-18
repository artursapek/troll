package market

func calculateEMA(interval MarketInterval, prevEMA float32, periods int) float32 {
  var multiplier float32 = 2 / (float32(periods) + 1)
  return (interval.CandleStick.Close - prevEMA) * multiplier + prevEMA
}

func (interval *MarketInterval) CalculateEMAs() {
  prev := (*interval).Prev()
  interval.EMA10 = calculateEMA(*interval, prev.EMA10, 10)
  interval.EMA21 = calculateEMA(*interval, prev.EMA21, 21)
}
