package main

import (
  "os"
  "fmt"
  "strconv"
  "github.com/pborman/getopt/v2"
  "github.com/yourbasic/graph"
  "github.com/lucasdc6/internet-topology/options"
  "github.com/lucasdc6/internet-topology/environment"
  "github.com/lucasdc6/internet-topology/bgpview/api"
)

func main() {
  debug := environment.GetDebugFor("main") || environment.GetDebug()
  // Declare options
  opAsInfo := getopt.BoolLong("as-info", 'a', "Show Asn info")
  optIx := getopt.BoolLong("ix", 'i', "Show IX connections")
  optHelp := getopt.BoolLong("help", 'h', "Show this help")
  opFull := getopt.BoolLong("full", 'f', "Show all the connections")

  // Parse
  getopt.Parse()

  if len(os.Args) < 2 || *optHelp {
    getopt.Usage()
    os.Exit(0)
  }

  if debug && *optIx {
    fmt.Println(*optIx)
  }

  if debug && *opFull {
    fmt.Println(*opFull)
  }

  if *opAsInfo {
    g := graph.New(99999)
    asn, err := strconv.ParseInt(getopt.Args()[0], 10, 64)
    if err != nil {
      fmt.Println("Error parsing as number")
      os.Exit(1)
    }
    fmt.Printf("Quering for ASN: %d\n", asn)
    as := api.GetAsnUpstreams(asn)
    for _, upstream := range as.Data.Ipv4Upstreams {
      options.CreateGraph(g, upstream.BgpPaths)
    }
    fmt.Println(g)
  }
}
