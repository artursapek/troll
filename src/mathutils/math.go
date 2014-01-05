package mathutils

import (
  "time"
)

// No generics. :(
func Max(nums ...float32) (max float32) {
  for i, num := range nums {
    if i == 0 {
      max = num
    } else if num > max {
      max = num
    }
  }
  return max
}

func Min(nums ...float32) (min float32) {
  for i, num := range nums {
    if i == 0 {
      min = num
    } else if num < min {
      min = num
    }
  }
  return min
}

func Abs(a float32) float32 {
  if a >= 0 {
    return a
  } else {
    return -a
  }
}

func Diff(a, b float32) float32 {
  return ((a - b) / ((a + b) / 2)) * 100
}

func RoundUpToNearestInterval(timestamp int64, timeInterval time.Duration) int64 {
  t := time.Unix(timestamp, 0)
  tRounded := t.Round(timeInterval * time.Second).Unix()

  if t.Unix() > tRounded {
    // That means we rounded down, and we wanted to round up
    tRounded += int64(timeInterval)
  }

  if tRounded%2 == 1 {
    tRounded += 1
  }

  return tRounded
}


