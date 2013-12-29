package ichimoku

import (
)

const tenkanPeriod = 8  // Traditionally 9
const kijunPeriod  = 11 // Traditionally 26
const chikouPeriod = kijunPeriod

// The main export:
type IchimokuLines struct {
  TenkenSen  float32
  KijunSen   float32
  ChikouSpan float32
  Kumo       struct {
    SenkouSpanA float32
    SenkouSpanB float32
  }
}

