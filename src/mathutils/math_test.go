package mathutils

import (
  "testing"
)

func TestMin(t *testing.T) {
  min := Min(20,40,10)
  if min != 10 {
    t.Errorf("MIN IS FUCKED")
  }
  max := Max(20,40,10)
  if max != 40 {
    t.Errorf("MAX IS FUCKED")
  }
  abs := Abs(-2)
  if abs != 2 {
    t.Errorf("ABS IS FUCKED")
  }
  abs2 := Abs(2)
  if abs2 != 2 {
    t.Errorf("ABS IS FUCKED")
  }

}
