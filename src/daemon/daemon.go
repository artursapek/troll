package daemon

import (
  "time"
)

type Daemon interface {
  Perform()
  Setup()
}

func Run (daemon Daemon, frequency time.Duration) {
  running := true
  daemon.Setup()
  for running {
    daemon.Perform()
    time.Sleep(frequency * 1000 * time.Millisecond)
  }
}

