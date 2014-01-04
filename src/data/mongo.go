package data

import (
  "env"
  "labix.org/v2/mgo"
)

const mongoUrl = "127.0.0.1"
const mongoDBName = "troll"

var mongoSession = newSession()

var Prices    *mgo.Collection
var Intervals *mgo.Collection
var Trades    *mgo.Collection

func newSession() *mgo.Session {
  session, err := mgo.Dial(mongoUrl)
  if err != nil {
    panic(err)
  }
  return session
}

func getPricesCollection() *mgo.Collection {
  // Frequent price check (15-30s)
  switch env.Env {
  case env.PRODUCTION:
    return getCollection("prices")
  case env.SIMULATION:
    return getCollection("test_prices")
  }
  return nil
}

func getIntervalsCollection() *mgo.Collection {
  // Candlestick and Ichimoku lines (2hr)
  switch env.Env {
  case env.PRODUCTION:
    return getCollection("intervals")
  case env.SIMULATION:
    return getCollection("test_2hrs")
  }
  return nil
}

func getTradesCollection() *mgo.Collection {
  // Candlestick and Ichimoku lines (2hr)
  switch env.Env {
  case env.PRODUCTION:
    return getCollection("trades")
  case env.SIMULATION:
    return getCollection("test_trades")
  }
  return nil
}

func getCollection(collection string) *mgo.Collection {
  return mongoSession.DB(mongoDBName).C(collection)
}

func init() {
  // Memoize collections
  Prices =    getPricesCollection()
  Intervals = getIntervalsCollection()
  Trades =    getTradesCollection()
}

