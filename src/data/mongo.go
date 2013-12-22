package data

import (
  "labix.org/v2/mgo"
)

const mongoUrl = "127.0.0.1"
const mongoDBName = "troll"

func NewSession() *mgo.Session {
  session, err := mgo.Dial(mongoUrl)
  if err != nil {
    panic(err)
  }
  return session
}

func GetCollection(collection string) *mgo.Collection {
  session := NewSession()
  return session.DB(mongoDBName).C(collection)
}

