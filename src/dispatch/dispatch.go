package dispatch

import (
  "fmt"
  "daemon"
  "time"
)

func Dispatch(command string) {
  switch command {
  case "prices":
    daemon.Run(daemon.PriceDaemon{}, time.Duration(15))
  case "status":
    daemon.Run(daemon.StatusDaemon{}, time.Duration(15))
  default:
    panic(fmt.Sprintf("Unknown command: %s", command))
  }
}
