package market

import (
  "time"
  "testing"
)

func TestRoundToNearest2Hour(t *testing.T) {
  // Test cases where it rounds up, and down

  UTC, _ := time.LoadLocation("UTC")

  time1 := time.Date(2014, time.January, 1, 4, 1, 2, 0, UTC).Unix()
  time1Expected := time.Date(2014, time.January, 1, 6, 0, 0, 0, UTC).Unix()
  time1Rounded := roundUpToNearest2Hour(time1)

  time2 := time.Date(2014, time.January, 1, 3, 58, 20, 0, UTC).Unix()
  time2Expected := time.Date(2014, time.January, 1, 4, 0, 0, 0, UTC).Unix()
  time2Rounded := roundUpToNearest2Hour(time2)


  if time1Rounded != time1Expected {
    t.Errorf("Rounded wrong. Expected: %d, got: %d", time1Expected, time1Rounded)
  }

  if time2Rounded != time2Expected {
    t.Errorf("Rounded wrong. Expected: %d, got: %d", time2Expected, time2Rounded)
  }
}

