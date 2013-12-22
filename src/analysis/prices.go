package analysis

import (
  "btce"
  "data"
  "labix.org/v2/mgo/bson"
)

var c = data.GetCollection("test_prices")

func GetTickersBetween(from, to int32) []btce.Ticker {
  var tickers []btce.Ticker
  c.Find(bson.M{"server_time": bson.M{"$gte": from, "$lt": to}}).All(&tickers)
  return tickers
}

func GetTickersBefore(timestamp int32, n int) []btce.Ticker {
  var tickers []btce.Ticker
  c.Find(bson.M{"server_time": bson.M{"$lte": timestamp}}).Sort("-server_time").Limit(n).All(&tickers)
  return tickers
}

type Tickers []btce.Ticker

func AverageDeviation(tickers Tickers) float32 {
  // A good measure of volatility
  amt := len(tickers)
  amtFloat := float32(amt)
  firstTicker := tickers[0]
  lastTicker  := tickers[amt - 1]
  slope := (lastTicker.Price - firstTicker.Price) / amtFloat

  var averageDeviation float32

  for i := 0; i < amt; i ++ {
    priceOnFlatGrowth := firstTicker.Price + (float32(i) * slope)
    dev := tickers[i].Price - priceOnFlatGrowth
    averageDeviation += dev / amtFloat
  }

  if averageDeviation < 0 { averageDeviation *= -1 } // abs

  return averageDeviation
}

func MinimumMaximum(tickers Tickers) (min, max float32) {
  amt := len(tickers)

  var price float32

  min = 9999999999999999999
  max = -9999999999999999999

  for i := 0; i < amt; i ++ {
    price = tickers[i].Price
    if price > max {
      max = price
    }
    if price < min {
      min = price
    }
  }

  return min, max
}

func AverageSlope(tickers Tickers) float32 {
  amt := len(tickers)
  amtFloat := float32(amt)

  var averageSlope float32

  for i := 0; i < amt - 1; i ++ {
    slope := tickers[i + 1].Price - tickers[i].Price
    averageSlope += slope / amtFloat
  }

  return averageSlope
}

func Percentile(tickers Tickers, ticker btce.Ticker) float32 {
  min, max := MinimumMaximum(tickers)
  diff := max - min
  return (ticker.Price - min) / diff
}

