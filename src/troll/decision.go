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


  if troll.Waiting() {

    if diff >= CLOSE_THRESHOLD {
      // R3
      if tenkan > kijun &&
         interval.CandleStick.Close > interval.SAR.Value {

        troll = troll.Buy(interval)
        fmt.Println("R3")
      }
    } else if diff >= OPEN_THRESHOLD {
      // R4
      if tenkan > kijun &&
         interval.CandleStick.Close > chikou &&
         mathutils.Min(tenkan, kijun) > mathutils.Max(senkouA, senkouB) {

        troll = troll.Buy(interval)
        fmt.Println("R4")
      }
    }


  } else {

    if diff >= CLOSE_THRESHOLD {
      // R1
      if tenkan < kijun &&
         interval.CandleStick.Close < interval.SAR.Value {

        troll = troll.Sell(interval)
        fmt.Println("R1")
      }
    } else if diff >= OPEN_THRESHOLD {
      // R2
      if tenkan < kijun &&
         interval.CandleStick.Close < chikou &&
         mathutils.Max(tenkan, kijun) > mathutils.Min(senkouA, senkouB) {

        troll = troll.Sell(interval)
        fmt.Println("R2")
      }
    }

  }

  return troll

  /*
  if diff >= CLOSE_THRESHOLD {
    prev := market.PrevInterval(interval)
    if (prev.Position == "long") && (tenkan < kijun) &&
       (interval.CandleStick.Close < interval.SAR.Value) {

      fmt.Println("R1")
      return troll.Sell(interval)

    } else if (prev.Position == "short") && (tenkan > kijun) &&
              (interval.CandleStick.Close > interval.SAR.Value) {

      fmt.Println("R2")
      return troll.Buy(interval)
    }
  } else if diff >= OPEN_THRESHOLD {
    if (tenkan > kijun) && (interval.CandleStick.Close > chikou) &&
       (mathutils.Min(tenkan, kijun) > mathutils.Max(senkouA, senkouB)) {

      fmt.Println("R3")
      market.SetPosition(interval, "long")
      return troll.Buy(interval)

    } else if (tenkan < kijun) && (interval.CandleStick.Close < chikou) &&
       (mathutils.Max(tenkan, kijun) > mathutils.Min(senkouA, senkouB)) {

      fmt.Println("R4")
      market.SetPosition(interval, "short")
      return troll.Sell(interval)
    }
  }

  return troll
  */
}
