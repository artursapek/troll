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

func printStatus() {
  var statuses []analysis.MarketStatus
  collectionStatuses.Find(nil).Limit(1).Sort("-servertime").All(&statuses)
  status := statuses[0]
  time := status.LocalTime
  secondsAgo := now() - time
  price := status.Price
  percentile := status.Analysis.Percentile

  fmt.Printf(fmt.Sprintf("%s%d seconds ago ", CLR_GREY, secondsAgo))
  fmt.Printf(fmt.Sprintf("%s$%.4f ", CLR_GREEN, price))

  fmt.Printf(fmt.Sprintf("%sPercentile: ", CLR_GREY))

  printP(percentile["6"])
  printP(percentile["12"])
  printP(percentile["24"])

  fmt.Printf(fmt.Sprintf("%sSlope: ", CLR_GREY))

  printS(status.Analysis.Slope["5"])
  printS(status.Analysis.Slope["10"])
  printS(status.Analysis.Slope["30"])
  printS(status.Analysis.Slope["60"])

  fmt.Printf(CLR_WHITE) // Reset
  fmt.Printf("         \r") // clear old shit
}

func (troll StatusDaemon) Setup() {
  header := "Troll"
  fmt.Println(header)
}

func (troll StatusDaemon) Perform() {
  printStatus()
}
