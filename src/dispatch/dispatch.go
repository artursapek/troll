// Entry point from main

package dispatch

import (
  "env"
  "fmt"
  "troll"
  "simulate"
  "time"
  "monitor"
)

// Use -live flag to run troll in live mode

func Dispatch(command string) {
  // Kill old process running this command
  killCommand(command)
  // Record new PID for this process
  writePID(command)

  switch command {
  case "run":
    if env.Prod() {
      troll.Run(troll.Troll{}, time.Duration(15))
    } else {
      simulate.Trade()
    }
  case "rebuild":
    troll.RebuildIntervals()
  case "status":
    troll.LastUpdate()
  case "http":
    monitor.StartServer()
  default:
    panic(fmt.Sprintf("Unknown command: %s", command))
  }
}
