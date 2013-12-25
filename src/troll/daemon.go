package troll

import (
  "time"
)

type Daemon interface {
  Perform() time.Duration
  Setup()
}

func Run (daemon Daemon, frequency time.Duration) {
  running := true
  daemon.Setup()
  for running {
    // A daemon can modify its own frequency via return value
    frequency = daemon.Perform()
    time.Sleep(frequency * 1000 * time.Millisecond)
  }
}

