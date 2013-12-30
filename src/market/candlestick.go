package market

type CandleStick struct {
  Open, Close, High, Low  float32
}

func createCandleStick(prices []MarketPrice) CandleStick {

  var low  float32 = 99999999999
  var high float32 = 0
  var open, close float32

  amt := len(prices)

  for i, price := range prices {
    if i == 0 {
      open = price.Price
    } else if i == amt - 1 {
      close = price.Price
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

