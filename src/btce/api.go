package btce

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

const publicApiUrl = "https://btc-e.com/api/2/btc_usd/%s"

func PublicApiRequest(action string) []byte {
  requestUrl := fmt.Sprintf(publicApiUrl, action)
  res, err := http.Get(requestUrl)
  if err != nil {
    panic(err)
  }
  body, _ := ioutil.ReadAll(res.Body)
  return body
}


