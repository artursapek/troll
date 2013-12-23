package analysis

import (
  "strconv"
  "labix.org/v2/mgo/bson"
)

func alltimePercentile(status MarketStatus) float32 {
  query := bson.M{
    "price": bson.M{ "$gte": status.Price },
    "servertime": bson.M{ "$lt": status.ServerTime },
  }
  ctHigher, _ := statusesCollection.Find(query).Count()
  ctAll,    _ := statusesCollection.Count()
  
  return float32(1.0) - (float32(ctHigher) / float32(ctAll))
}

func calculatePercentileMap(status MarketStatus) Metrics {
  metrics := make(Metrics)
  for i := 0; i < 5; i ++ {
    hrs := hourlyMetrics[i]
    hrsString := strconv.Itoa(hrs)
    r := status.Analysis.Range[hrsString]
    d := r.Max - r.Min
    perc := (status.Price - r.Min) / d
    if perc < 0 {
      perc *= -1
    }
    metrics[hrsString] = perc
  }
  metrics["all"] = alltimePercentile(status)
  return metrics
}


