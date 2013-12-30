package simulate

import (
  "fmt"
  "os"
  "data"
  "btce"
  "market"
  "troll"
  "strconv"
)

// Start at any point in the simulation
func MakeTrollFromStatus(status market.MarketPrice) troll.Troll {
  funds := troll.FundsStatus{
    BTC: 1.0,
    USD: 0,
  }
  return troll.Troll{ 
    Funds: funds,
    LastTrade: btce.OwnTrade{
      Pair: "btc_usd",
      Type: "buy",
      Amount: 1.0,
      Rate: status.Price,
      Timestamp: status.Time.Server,
    },
  }

}

func init() {
  // Clear the intervals, which we're generating ourselves.
  data.Intervals.DropCollection()
}

func Simulate() {
  var skip, limit int
  if len(os.Args) < 3 {
    skip = 0
    limit = 999999999999
  } else {
    skip, _ = strconv.Atoi(os.Args[2])
    limit, _ = strconv.Atoi(os.Args[3])
  }

  var prices []market.MarketPrice
  data.Prices.Find(nil).Skip(skip).Limit(limit).All(&prices)

  amt := len(prices)

  //self := MakeTrollFromStatus(prices[0])

  for i := 0; i < amt; i ++ {
    price := prices[i]
    //now := price.Time.Local

    market.ProcessPrice(price)
    fmt.Printf(".")
  }
}

