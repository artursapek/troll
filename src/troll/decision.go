package troll

import (
  "mathutils"
  "fmt"
  "market"
)

// Thresholds, configurable
const OPEN_THRESHOLD float32 =  0.10
const CLOSE_THRESHOLD float32 = 2.18

// Troll is great at decision making
func (troll Troll) Decide(interval market.MarketInterval) Troll {
//  fmt.Printf("%d,%f,%f\n", interval.Time.Close, interval.CandleStick.Open, interval.CandleStick.Close)

  tenkan, kijun, senkouA, senkouB, chikou := market.UnpackIchimoku(interval.Ichimoku)

  //fmt.Println(tenkan, kijun, senkouA, senkouB, chikou)

  diff := ((tenkan - kijun) / ((tenkan + kijun) / 2)) * 100

  diff = mathutils.Abs(diff)

  // Is the SAR below the price?
  sarBullish := interval.CandleStick.Close > interval.SAR.Value 

  maxKumo := mathutils.Max(senkouA, senkouB)
  minKumo := mathutils.Min(senkouA, senkouB)

  price := interval.CandleStick.Close

  tkDiff := mathutils.Diff(tenkan, kijun)

  chikouSpan := price - chikou

  if troll.Waiting() {

    if diff >= CLOSE_THRESHOLD {
      // R3
      // Signs point to bullish
      if tenkan > kijun && 
         tkDiff > 2.4 &&
         sarBullish {

        troll = troll.Buy(interval)
        fmt.Println("R3")
      }
    } else if diff >= OPEN_THRESHOLD {
      // R4
      // Bullish
      if tenkan > kijun &&
         chikouSpan <= 0 &&
         mathutils.Min(tenkan, kijun) > maxKumo {
         //mathutils.Min(tenkan, kijun) > mathutils.Max(senkouA, senkouB) {

        troll = troll.Buy(interval)
        fmt.Println("R4")
      }
    }

  } else {

    if diff >= CLOSE_THRESHOLD {
      // R1
      if tenkan < kijun &&
         tkDiff <= 2.1 &&
         !sarBullish {

        troll = troll.Sell(interval)
        fmt.Println("R1")
      }
    } else if diff >= OPEN_THRESHOLD {
      // R2
      if tenkan < kijun &&
         chikouSpan >= 0 &&
         mathutils.Max(tenkan, kijun) < minKumo {

        troll = troll.Sell(interval)
        fmt.Println("R2")
      }
    }

  }
  return troll
}
