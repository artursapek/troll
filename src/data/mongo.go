package data

import (
  "env"
  "labix.org/v2/mgo"
)

const mongoUrl = "127.0.0.1"
const mongoDBName = "troll"

var mongoSession = NewSession()

func NewSession() *mgo.Session {
  session, err := mgo.Dial(mongoUrl)
  if err != nil {
    panic(err)
  }
  return session
}

func GetStatusCollection() *mgo.Collection {
  if env.Env == "production" {
    return mongoSession.DB(mongoDBName).C("statuses")
  } else if env.Env == "simulation" {
    return mongoSession.DB(mongoDBName).C("test_prices")
  }
  return nil
}

func GetCollection(collection string) *mgo.Collection {
  return mongoSession.DB(mongoDBName).C(collection)
}
