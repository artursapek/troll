package dispatch

import (
  "fmt"
  "troll"
  "simulate"
  "time"
)

func Dispatch(command string) {
  killCommand(command)
  writePID(command)
  switch command {
  case "run":
    troll.Run(troll.Troll{}, time.Duration(30))
  case "status":
    troll.Run(troll.StatusDaemon{}, time.Duration(1))
  case "runsim":
    simulate.Simulate()
  default:
    panic(fmt.Sprintf("Unknown command: %s", command))
  }
}
