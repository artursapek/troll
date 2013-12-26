// Entry point from main

package dispatch

import (
  "env"
  "fmt"
  "troll"
  "simulate"
  "time"
)

// Use -live flag to run troll in live mode

func Dispatch(command string) {
  // Kill old process running this command
  killCommand(command)
  // Record new PID for this process
  writePID(command)

  if env.Env == "production" {
    // PRODUCTION
    switch command {
    case "run":
      troll.Run(troll.Troll{}, time.Duration(15))
    case "status":
      troll.Run(troll.StatusDaemon{}, time.Duration(1))
    default:
      panic(fmt.Sprintf("Unknown command: %s", command))
    }

  } else {
    // SIMULATION
    switch command {
    case "run":
      simulate.Simulate()
    default:
      panic(fmt.Sprintf("Unknown command: %s", command))
    }
  }
}
