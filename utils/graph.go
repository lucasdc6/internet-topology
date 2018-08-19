package utils

import (
  "fmt"
  "strconv"
  "strings"
  "github.com/yourbasic/graph"
  "github.com/lucasdc6/internet-topology/environment"
)

func GraphToJson(g *graph.Mutable, from int) string {
  var str strings.Builder
  graph.BFS(g, from, func(v, w int, _ int64) {
    fmt.Fprintf(&str, "%d,%d\n", v, w)
  })
  return str.String()
}

func AddToGraph(g *graph.Mutable, paths []string, deepLevel int) {
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
      g.Add(actual, asni)
      actual = asni
    }
  }
  if debugEnv {
    fmt.Println(g)
  }
}
