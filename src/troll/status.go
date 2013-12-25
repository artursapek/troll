package troll

import (
  "data"
  "analysis"
  "fmt"
  "time"
)

const CLR_WHITE  = "\x1b[37;1m"
const CLR_GREY   = "\x1b[30;1m"
const CLR_GREEN  = "\x1b[32;1m"
const CLR_YELLOW = "\x1b[33;1m"
const CLR_RED    = "\x1b[31;1m"

var collectionStatuses = data.GetCollection("statuses")

type StatusDaemon struct{}

func now() int32 {
  return int32(time.Now().Unix())
}

func printP(percentile float32) {
  var color string
  if percentile < 0.33 {
    color = CLR_RED
  } else if percentile < 0.66 {
    color = CLR_YELLOW
  } else {
    color = CLR_GREEN
  }
  fmt.Printf(fmt.Sprintf("%s%.2f", color, percentile * 100))
  fmt.Printf("%% ")
}

func printS(slope float32) {
  var color string
  if slope > -5 && slope < 5 {
    color = CLR_YELLOW
  } else if slope <= -5 {
    color = CLR_RED
  } else {
    color = CLR_GREEN
  }
  if slope < 0 {
    fmt.Printf(fmt.Sprintf("%s-$%.2f ", color, slope * -1))
  } else {
    fmt.Printf(fmt.Sprintf("%s+$%.2f ", color, slope))
  }
}

func clear() {
  // Clear the screen
  fmt.Printf("\033c")
}

func heading(head string) {
  fmt.Printf(fmt.Sprintf("\n\n  %s%-12s", CLR_GREY, head))
}

func printStatus() {
  clear()

  // Get newest status
  var statuses []analysis.MarketStatus
  collectionStatuses.Find(nil).Limit(1).Sort("-servertime").All(&statuses)
  status := statuses[0]

  fmt.Println("")

  price := status.Price
  fmt.Printf(fmt.Sprintf("  %s$%.4f", CLR_GREEN, price))

  lastUpdateTime := status.LocalTime
  secondsAgo := now() - lastUpdateTime
  heading(fmt.Sprintf("%ds ago", secondsAgo))

  percentile := status.Analysis.Percentile

  heading("Percentile")

  printP(percentile["6"])
  printP(percentile["12"])
  printP(percentile["24"])

  heading("Slope")

  printS(status.Analysis.Slope["5"])
  printS(status.Analysis.Slope["10"])
  printS(status.Analysis.Slope["30"])
  printS(status.Analysis.Slope["60"])

  fmt.Printf(CLR_WHITE) // Reset
  fmt.Println("") // clear old shit
}

func (troll StatusDaemon) Setup() {
  header := "Troll"
  fmt.Println(header)
}

func (troll StatusDaemon) Perform() time.Duration {
  printStatus()
  return 1 // Never change the interval from 1 second
}
