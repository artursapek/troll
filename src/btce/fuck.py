# If you find this sample useful, please feel free to donate :)
# LTC: LePiC6JKohb7w6PdFL2KDV1VoZJPFwqXgY
# BTC: 1BzHpzqEVKjDQNCqV67Ju4dYL68aR8jTEe
 
import httplib
import urllib
import json
import hashlib
import hmac
 
# Replace these with your own API key data
BTC_api_key =    "ES624GXB-98HIGRUB-8LU9HZS5-DLRB1OZ7-N32I2YL3"
BTC_api_secret = "9b64f25f7659bde648d0685e3283476fc1eac9b26bb94b5dbc1a42f438dd6580"
FUUCK = 'cc4fd37271a497193350927dcb0e98f58189e67947b7659b87934a0bb0f3125b75ff6c4d0f60ba32411a9d80c0a9445e4ef6a75a7af113809a528d72e6c9f27a'

# Come up with your own method for choosing an incrementing nonce
nonce = 13
 
# method name and nonce go into the POST parameters
params = {"method":"getInfo",
          "nonce": nonce}
params = urllib.urlencode(params)
print params
 
# Hash the params string to produce the Sign header value
H = hmac.new(BTC_api_secret, digestmod=hashlib.sha512)
H.update(params)
sign = H.hexdigest()
 
print sign
headers = {"Content-type": "application/x-www-form-urlencoded",
                   "Key":BTC_api_key,
                   "Sign":sign}
conn = httplib.HTTPSConnection("btc-e.com")
conn.request("POST", "/tapi", params, headers)
response = conn.getresponse()
 
print response.status, response.reason
print json.load(response)
 
conn.close()
