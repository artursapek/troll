package daemon

import (
  "time"
)

type Daemon interface {
  Perform()
}

func Run (daemon Daemon, frequency time.Duration) {
  running := true
  for running {
    daemon.Perform()
    time.Sleep(frequency * 1000 * time.Millisecond)
  }
}

