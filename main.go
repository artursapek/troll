package main

import (
  "os"
  "dispatch"
)

func main () {
  dispatch.Dispatch(os.Args[1])
}
