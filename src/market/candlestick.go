package market

type CandleStick struct {
  Open, Close, High, Low  float32
}

func createCandleStick(prices []MarketPrice) CandleStick {

  var low  float32
  var high float32
  var open, close float32

  amt := len(prices)

  for i, price := range prices {
    if i == 0 {
      close = price.Price

      // Default to first price, then do the checks
      low = price.Price
      high = price.Price

    } else if i == amt - 1 {
      open = price.Price
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

