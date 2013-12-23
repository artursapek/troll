package analysis

import (
  "fmt"
  "labix.org/v2/mgo/bson"
)

func statusesFromPastNHours (status MarketStatus, n int32) (statuses []MarketStatus) {
  // Go back in time n hours in seconds
  start := status.ServerTime - (n * 60 * 60)
  // Find all statuses between the two dates, and unpack them into statuses variable
  query := bson.M{ "servertime": bson.M{ "$gte": start, "$lt": status.ServerTime }}
  err := statusesCollection.Find(query).Sort("-servertime").All(&statuses)
  if err != nil {
    fmt.Println(err) // Log it, and return empty
  }
  return statuses
}

func filterPastNHours (statuses []MarketStatus, n, now int32) (results []MarketStatus) {
  start := now - (n * 60 * 60)
  for i := 0; i < len(statuses); i ++ {
    status := statuses[i]
    if status.ServerTime > start {
      results = append(results, status)
    } else {
      // They come from Mongo in order, so we can break
      // as soon as we leave the range we wanted.
      break
    }
  }
  return results
}

func filterPastNMinutes (statuses []MarketStatus, n, now int32) (results []MarketStatus) {
  start := now - (n * 60)
  for i := 0; i < len(statuses); i ++ {
    status := statuses[i]
    if status.ServerTime > start {
      results = append(results, status)
    } else {
      break
    }
  }
  return results
}

