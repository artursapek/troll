package market

// Calculating parabolic SAR (stop and reverse)

const SAR_ACC_INCREMENT float32 = 0.025
const SAR_ACC_MAX float32 = 0.15

type ParabolicSAR struct {
  Value float32
  // "long" or "short"
  Position string
  // Must keep track of the acceleration factor
  // for the next candlestick
  Acc float32
  AccD float32
}

// No generics. :(
func max(a, b float32) float32 {
  if (a > b) {
    return a
  } else {
    return b
  }
}

func min(a, b float32) float32 {
  if (a < b) {
    return a
  } else {
    return b
  }
}

func abs(a float32) float32 {
  if a >= 0 {
    return a
  } else {
    return -a
  }
}

func CalculateParabolicSAR(curr, prev, prevPrev MarketInterval) (SAR ParabolicSAR) {

  // Init acceleration factor
  SAR.Acc = prev.SAR.Acc
  SAR.Position = prev.SAR.Position

  var epCurr, epPrev float32

  if prev.SAR.Position == "long" {
    epCurr = max(curr.CandleStick.High, prev.CandleStick.High)
    epPrev = max(prev.CandleStick.High, prevPrev.CandleStick.High)

    // If we have a new extreme price increment the acceleration factor
    if epCurr > epPrev { 
      SAR.Acc += SAR_ACC_INCREMENT
      if SAR.Acc > SAR_ACC_MAX {
        SAR.Acc = SAR_ACC_MAX
      }
    }

  } else {
    epCurr = min(curr.CandleStick.Low, prev.CandleStick.Low)
    epPrev = min(prev.CandleStick.Low, prevPrev.CandleStick.Low)

    if epCurr < epPrev {
      SAR.Acc += SAR_ACC_INCREMENT
      if SAR.Acc > SAR_ACC_MAX {
        SAR.Acc = SAR_ACC_MAX
      }
    }
  }


  SAR.AccD = SAR.Acc * abs(epCurr - prev.SAR.Value)

  if prev.SAR.Position == "long" {
    if (prev.SAR.Value + prev.SAR.AccD) > curr.CandleStick.Low {
      // Switch & reverse
      SAR.Position = "short"
      SAR.Value = epPrev
      SAR.Acc = SAR_ACC_INCREMENT
    } else {
      SAR.Position = "long"
      SAR.Value = min(min((prev.SAR.Value + prev.SAR.AccD), prev.CandleStick.Low), prevPrev.CandleStick.Low)
    }
  } else {
    if (prev.SAR.Value - prev.SAR.AccD) < curr.CandleStick.High {
      // Switch & reverse
      SAR.Position = "long"
      SAR.Value = epPrev
      SAR.Acc = SAR_ACC_INCREMENT
    } else {
      SAR.Position = "short"
      SAR.Value = max(max((prev.SAR.Value - prev.SAR.AccD), prev.CandleStick.High), prevPrev.CandleStick.High)
    }
  }
  return SAR
}
