package env

import (
  "os"
)

const PRODUCTION string = "production"
const SIMULATION string = "simulation"

var Env string

func init() {
  // Operate in live mode using the -live flag

  if len(os.Args) == 3 && os.Args[2] == "-live" {
    Env = PRODUCTION
  } else {
    Env = SIMULATION
  }
}
