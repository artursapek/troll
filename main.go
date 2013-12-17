package main

import (
  "btce"
  "net/url"
  "io/ioutil"
)

func main () {
  values := url.Values{}
  res := btce.ApiRequest("getInfo", values)
  body, _ := ioutil.ReadAll(res.Body)
  print(string(body))
}
