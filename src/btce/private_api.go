package btce

import (
  "bytes"
  "net/http"
  "net/url"
  "time"
  "crypto/sha512"
  "crypto/hmac"
  "strconv"
  "encoding/hex"
)

const API_KEY string =    "ES624GXB-98HIGRUB-8LU9HZS5-DLRB1OZ7-N32I2YL3"
const API_SECRET string = "9b64f25f7659bde648d0685e3283476fc1eac9b26bb94b5dbc1a42f438dd6580"

func nonce() string {
  return strconv.FormatInt(int64(time.Now().Unix()), 10)
}

func hash(params string) string {
  hash := sha512.New
  h := hmac.New(hash, []byte(API_SECRET))
  h.Write([]byte(params))
  return hex.EncodeToString(h.Sum(nil))
}

func ApiRequest(action string, params url.Values) *http.Response {
  client := &http.Client{}

  params.Set("method", action)
  params.Set("nonce", nonce())

  paramsEncoded := params.Encode()
  body := bytes.NewBufferString(paramsEncoded)
  req, _ := http.NewRequest("POST", "https://btc-e.com/tapi", body)

  req.Header.Set("Content-type", "application/x-www-form-urlencoded")
  req.Header.Set("Key", API_KEY)
  req.Header.Set("Sign", hash(paramsEncoded))

  response, err := client.Do(req)
  if err != nil {
    panic(err)
  } else {
    return response
  }
}

