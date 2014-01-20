package monitor

import (
  "fmt"
  "net/http"
  "data"
  "market"
  "encoding/json"
  "btce"
  "time"
)

var cachedIntervals market.MarketIntervals
var cacheExpiration int64

func getIntervals() market.MarketIntervals {
  now := time.Now().Unix()

  fmt.Println(now, cacheExpiration)

  if now > cacheExpiration {
    // Cache is invalid. Run the query again.
    // Initially cacheExpiration will be 0,
    // so this will always run the first time
    var intervals market.MarketIntervals
    var limit int = 12 * 30 * 3 // 3 mos
    data.Intervals.Find(nil).Sort("-time.close").Limit(limit).All(&intervals)
    cacheExpiration = intervals[0].Time.Close + 60 * 60 * 2 // Reset the expiration time
    cachedIntervals = intervals
    fmt.Println("Ran interval query")
    return intervals
  } else {
    fmt.Println("Used cached intervals")
    return cachedIntervals
  }
}

type IntervalsResponse struct {
  Intervals market.MarketIntervals
  PingIn int64
}

func pingIn(interval market.MarketInterval) int64 {
  now := time.Now().Unix()
  return market.INTERVAL_PERIOD - (now - interval.Time.Close) + 60 * 2
}

func intervalsHandler(rw http.ResponseWriter, req *http.Request) {

  response := IntervalsResponse{}
  response.Intervals = getIntervals()
  // Tell client to ping again when the next interval
  // is ready, plus two minutes.
  response.PingIn = pingIn(response.Intervals[0])

  body, err := json.Marshal(response)
  if err != nil {
    panic(err)
  } else {
    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(rw, string(body))
  }
}

func latestIntervalHandler(rw http.ResponseWriter, req *http.Request) {
  latestInterval := market.LastInterval()

  response := IntervalsResponse{}
  response.Intervals = []market.MarketInterval{latestInterval}
  response.PingIn = pingIn(latestInterval)

  body, err := json.Marshal(response)
  if err != nil {
    panic(err)
  } else {
    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(rw, string(body))
  }

}

func tradesHandler(rw http.ResponseWriter, req *http.Request) {
  var trades []btce.OwnTrade
  data.Trades.Find(nil).All(&trades)
  body, _ := json.Marshal(trades)

  rw.Header().Set("Access-Control-Allow-Origin", "*")
  rw.Header().Set("Content-Type", "application/json")
  fmt.Fprintf(rw, string(body))
}

func StartServer() {
  getIntervals() // Cache them for the first time
  http.HandleFunc("/intervals.json", intervalsHandler)
  http.HandleFunc("/latest-interval.json", latestIntervalHandler)
  http.HandleFunc("/trades.json", tradesHandler)
  http.ListenAndServe(":8001", nil)
}
