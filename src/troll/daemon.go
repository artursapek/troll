package troll

import (
  "time"
)

type Daemon interface {
  Perform()
  Setup()
}

func Run (troll Daemon, frequency time.Duration) {
  running := true
  troll.Setup()
  for running {
    troll.Perform()
    time.Sleep(frequency * 1000 * time.Millisecond)
  }
}

