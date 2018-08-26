package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "github.com/pborman/getopt/v2"
  "gonum.org/v1/gonum/graph"
  "gonum.org/v1/gonum/graph/simple"
  "gonum.org/v1/gonum/graph/encoding/dot"
  "gonum.org/v1/gonum/graph/traverse"
  "github.com/lucasdc6/internet-topology/utils"
  "github.com/lucasdc6/internet-topology/environment"
  "github.com/lucasdc6/internet-topology/bgpview/api"
)

func main() {
  debug := environment.GetDebugFor("main") || environment.GetDebug()
  // Declare options
  optPipeMode := getopt.BoolLong("pipe", 0, "Set pipe mode.")
  optAsn := getopt.IntLong("asn", 0, -1, "Set AS number")
  optIx := getopt.BoolLong("ix", 'i', "Show IX connections")
  optPeers := getopt.BoolLong("peers", 'p', "Show Peers connections")
  optHelp := getopt.BoolLong("help", 'h', "Show this help")
  optFull := getopt.BoolLong("full", 'f', "Show all the connections")
  optDeepLevel := getopt.IntLong("deep", 0, 99, "Set the deep level. Default to full path")
  optOutput := getopt.StringLong("output", 'o', "/tmp/internet-topology", "Set the output file. Default to /tmp/internet-topology")

  // Parse
  getopt.Parse()

  // Check args
  if len(os.Args) < 2 || *optHelp {
    getopt.Usage()
    os.Exit(0)
  }

  if debug && *optIx {
    fmt.Println(*optIx)
  }

  if debug && *optFull {
    fmt.Println(*optFull)
  }

  // Init graph
  g := simple.NewDirectedGraph()
  if *optAsn > 0 {
    if !*optPipeMode {
      fmt.Printf("Quering for ASN: %d\n", *optAsn)
    }
    as := api.GetAsnUpstreams(*optAsn)
    for _, upstream := range as.Data.Ipv4Upstreams {
      utils.AddToGraph(g, upstream.BgpPaths, *optDeepLevel+1)
    }
    if *optFull {
      b := traverse.BreadthFirst{
        Visit: func(u, v graph.Node) {
          if v.ID() != int64(*optAsn) {
            if !*optPipeMode {
              fmt.Printf("Quering for ASN: %d\n", v)
            }
            as := api.GetAsnUpstreams(int(v.ID()))
            for _, upstream := range as.Data.Ipv4Upstreams {
              utils.AddToGraph(g, upstream.BgpPaths, *optDeepLevel+1)
            }
          }
        }}
      b.Walk(g, g.Node(int64(*optAsn)), nil)
    }
  }
  if *optIx {
    if !*optPipeMode {
      fmt.Printf("Quering for ASN IX: %d\n", *optAsn)
    }
    asIxs := api.GetAsnIxs(*optAsn)
    for _, asIx := range asIxs.Data {
      if !*optPipeMode {
        fmt.Printf("Quering for IX (%s) members: %d\n", asIx.Name, asIx.IxId)
      }
      ix := api.GetIx(asIx.IxId)
      for _, member := range ix.Data.Members {
        if member.Asn != *optAsn {
          g.SetEdge(simple.Edge{F: simple.Node(*optAsn), T: simple.Node(member.Asn)})
        }
      }
    }
  }

  if *optPeers {
    if !*optPipeMode {
      fmt.Printf("Quering for ASN Peers: %d\n", *optAsn)
    }
    asPeers := api.GetAsnPeers(*optAsn)
    for _,asPeer := range asPeers.Data.Ipv4Peers {
      g.SetEdge(simple.Edge{F: simple.Node(*optAsn), T: simple.Node(asPeer.Asn)})
    }
  }
  if *optPipeMode {
    fmt.Printf("%s", utils.GraphToData(g, *optAsn))
  } else {
    fmt.Printf("Saving file in %s\n", *optOutput)
  }

  str, _ := dot.Marshal(g, "", "", "  ", false)
  ioutil.WriteFile(fmt.Sprintf("%s.dot", *optOutput), str, 0644)
  ioutil.WriteFile(fmt.Sprintf("%s.data", *optOutput), []byte(utils.GraphToData(g, *optAsn)), 0644)
}
