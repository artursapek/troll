package analysis

import (
  "strconv"
  "errors"
  "labix.org/v2/mgo/bson"
)

func PreviousStatus(status MarketStatus) (MarketStatus, error) {
  var unpack []MarketStatus
  query := bson.M{ "servertime": bson.M{ "$lt": status.ServerTime }}
  statusesCollection.Find(query).Sort("-servertime").Limit(1).All(&unpack)

  if len(unpack) == 0 {
    return MarketStatus{}, errors.New("No statuses found")
  }

  return unpack[0], nil
}

func calculateEMA(status MarketStatus, timePeriod int) float32 {
  prev, err := PreviousStatus(status)
  if err != nil {
    return 1
  }

  prevEMA := prev.Analysis.EMA[strconv.Itoa(timePeriod)]

  if prevEMA <= 0 {
    prevEMA = float32(1)
  }

  var m float32 = (2 / (float32(timePeriod) + 1))

  return ((status.Price * m) + (prevEMA * (1 - m)))
}

