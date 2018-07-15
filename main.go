package main

import (
  "os"
  "fmt"
  "github.com/pborman/getopt/v2"
)

func main() {
  // Declare options
  optIx := getopt.BoolLong("ix", 'i', "Show IX connections")
  optHelp := getopt.BoolLong("help", 'h', "Show this help")
  opFull := getopt.BoolLong("full", 'f', "Show all the connections")

  // Parse
  getopt.Parse()

  if *optHelp {
    getopt.Usage()
    os.Exit(0)
  }

  if *optIx {
    fmt.Println(*optIx)
  }

  if *opFull {
    fmt.Println(*opFull)
  }
}
