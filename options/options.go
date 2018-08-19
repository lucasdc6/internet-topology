package options

import (
  "fmt"
  "strconv"
  "strings"
  "github.com/yourbasic/graph"
  "github.com/lucasdc6/internet-topology/environment"
)

func CreateGraph(g *graph.Mutable, paths []string) {
  debugEnv := environment.GetDebugFor("options") || environment.GetDebug()
  for _, path := range paths {
    var actual, asni int
    asns := strings.Split(path, " ")
    actuals, asns := asns[0], asns[1:]
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
