package utils

import (
  "fmt"
  "strconv"
  "strings"
  graph "gonum.org/v1/gonum/graph/simple"
  "github.com/lucasdc6/internet-topology/environment"
)

// GraphToData - Convert graph to str
func GraphToData(g *graph.DirectedGraph, from int) string {
  var str strings.Builder
  for _, edge := range g.Edges() {
    fmt.Fprintf(&str, "%d,%d\n", edge.From(), edge.To())
  }
  return str.String()
}

// AddToGraph - Add paths array with deepLevel to graph g
func AddToGraph(g *graph.DirectedGraph, paths []string, deepLevel int) {
  debugEnv := environment.GetDebugFor("options") || environment.GetDebug()
  for _, path := range paths {
    var high, actual, asni int
    asns := strings.Split(path, " ")

    if deepLevel < len(asns) && deepLevel != 0 {
      high = deepLevel
    } else {
      high = len(asns)
    }

    actuals, asns := asns[0], asns[1:high]
    actual, _ = strconv.Atoi(actuals)
    if debugEnv {
      fmt.Printf("Actual path: %s\n", path)
    }
    for _, asn := range asns {
      asni, _ = strconv.Atoi(asn)
      if debugEnv {
        fmt.Printf("Added %d -> %d to graph\n\n", actual, asni)
      }
      g.SetEdge(graph.Edge{F: graph.Node(actual), T: graph.Node(asni)})
      actual = asni
    }
  }
  if debugEnv {
    fmt.Println(g)
  }
}
