package monitor

import (
  "fmt"
  "net/http"
  "data"
  "market"
  "encoding/json"
  "btce"
  "time"
  "net/url"
  "strconv"
  "labix.org/v2/mgo/bson"
)

var cachedIntervals market.MarketIntervals
var cacheExpiration int64

func getIntervals(after int64) market.MarketIntervals {
  // Given closing time "after"
  now := time.Now().Unix()

  fmt.Println(now, cacheExpiration)

  // Cache is invalid. Run the query again.
  // Initially cacheExpiration will be 0,
  // so this will always run the first time
  var intervals market.MarketIntervals
  var limit int = 12 * 30 * 3 // 3 mos
  data.Intervals.Find(bson.M{"time.close": bson.M{"$gt": after}}).Sort("-time.close").Limit(limit).All(&intervals)
  if len(intervals) == 0 {
    return intervals
  }
  fmt.Println("Ran interval query")
  return intervals
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
  query, _ := url.ParseQuery(req.URL.RawQuery)

  after, _ := strconv.Atoi(query["after"][0])

  response := IntervalsResponse{}
  response.Intervals = getIntervals(int64(after))
  // Tell client to ping again when the next interval
  // is ready, plus two minutes.
  if len(response.Intervals) == 0 {
    response.PingIn = market.INTERVAL_PERIOD - (time.Now().Unix() - int64(after)) + 60 * 2
  } else {
    response.PingIn = pingIn(response.Intervals[0])

  }


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
  http.HandleFunc("/intervals.json", intervalsHandler)
  http.HandleFunc("/latest-interval.json", latestIntervalHandler)
  http.HandleFunc("/trades.json", tradesHandler)
  http.ListenAndServe(":8001", nil)
}
