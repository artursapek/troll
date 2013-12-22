package dispatch

import (
  "fmt"
  "daemon"
  "simulate"
  "time"
)

func Dispatch(command string) {
  killCommand(command)
  writePID(command)
  switch command {
  case "run":
    daemon.Run(daemon.TrollDaemon{}, time.Duration(15))
  case "status":
    daemon.Run(daemon.StatusDaemon{}, time.Duration(15))
  case "runsim":
    simulate.Iterate()
  default:
    panic(fmt.Sprintf("Unknown command: %s", command))
  }
}
