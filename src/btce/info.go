package btce

import (
  "encoding/json"
  "fmt"
  "net/url"
)

type Funds map[string]float32

type Info struct {
  Funds Funds
}

type InfoResponse struct {
  Success int
  Return  Info
  Error   string
}

func decodeInfo(body []byte) InfoResponse {
  var info InfoResponse
  json.Unmarshal(body, &info)
  if info.Success != 1 {
    fmt.Println(info.Error)
  }
  return info
}

func getInfo() Info {
  responseBody := ApiRequest("getInfo", url.Values{})
  return decodeInfo(responseBody).Return
}

func GetFunds() Funds {
  info := getInfo()
  return info.Funds
}
