package env

import (
  "os"
)

var Env string

func init() {
  // Operate in live mode using the -live flag

  if len(os.Args) == 3 && os.Args[2] == "-live" {
    Env = "production"
    return
  }
  Env = "simulation"
}
