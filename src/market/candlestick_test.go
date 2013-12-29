package market

import (
  "testing"
)

func TestCandleStick(t *testing.T) {
  var prices = []MarketPrice{
    MarketPrice{ Price: 10 },
    MarketPrice{ Price: 15 },
    MarketPrice{ Price: 20 },
    MarketPrice{ Price: 15 },
    MarketPrice{ Price: 8 },
    MarketPrice{ Price: 12 },
  }

  candleStick := CreateCandleStick(prices)

  if candleStick.Open != 10 {
    t.Errorf("Open price incorrect: %f", candleStick.Open)
  }
  if candleStick.Close != 12 {
    t.Errorf("Close price incorrect: %f", candleStick.Close)
  }
  if candleStick.High != 20 {
    t.Errorf("High price incorrect: %f", candleStick.High)
  }
  if candleStick.Low != 8 {
    t.Errorf("Low price incorrect: %f", candleStick.Low)
  }

}
