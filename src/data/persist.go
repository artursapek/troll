package data

import (
  "os"
  "fmt"
  "bufio"
  "labix.org/v2/mgo"
)

const dataDir = "src/troll/memory/data/%s.data"

func fn(name string) string {
  return fmt.Sprintf(dataDir, name)
}

func Set (key, value string) {
  print(value)
}

func Get (key string) {

}

func Append(key, identifier, lines string) {
  filename := fn(key)
  f, _ := os.OpenFile(filename, 0666, os.ModeAppend)
  defer f.Close()
  reader := bufio.NewReader(f)
  line, _ := reader.ReadString('\n')
  // It's ASCII, so it's safe to assume it's one byte per char
  lineLength := len(line)

  stat, err := os.Stat(filename)

  if err != nil {
    panic(err)
  }

  buf := make([]byte, lineLength)

  f.ReadAt(buf, stat.Size() - int64(lineLength - 1))
  print(string(buf))
}




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

