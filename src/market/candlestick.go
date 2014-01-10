package market

type CandleStick struct {
  Open, Close, High, Low  float32
}

func createCandleStick(prices []MarketPrice, lastClose float32) CandleStick {

  var low  float32
  var high float32
  var open, close float32

  // Line up open with last close
  open = lastClose

  for i, price := range prices {
    if i == 0 {
      close = price.Price

      // Default to first price, then do the checks
      low = price.Price
      high = price.Price
    }
    if price.Price > high {
      high = price.Price
    }
    if price.Price < low {
      low = price.Price
    }
  }

  return CandleStick{
    Open:  open,
    Close: close,
    High:  high,
    Low:   low,
  }
}

