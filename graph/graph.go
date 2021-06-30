package graph

import (
	"fmt"
	"strings"
	"time"

	"github.com/lucasdc6/internet-topology/bgpview/api"
	"github.com/lucasdc6/internet-topology/bgpview/types"
	"github.com/lucasdc6/internet-topology/environment"
	graph "gonum.org/v1/gonum/graph/simple"
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
func AddToGraph(g *graph.DirectedGraph, peers []types.AsnUpDownstream, from, deepLevel int) {
	if deepLevel == 0 {
		return
	}
	debugEnv := environment.GetDebugFor("options") || environment.GetDebug()
	for _, peer := range peers {
		if g.Node(int64(peer.Asn)) != nil {
			continue
		}
		g.SetEdge(graph.Edge{F: graph.Node(from), T: graph.Node(peer.Asn)})
		if deepLevel != 0 {
			fmt.Printf("Quering for ASN: %d\n", peer.Asn)
			time.Sleep(1 * time.Second)
			as := api.GetAsnUpstreams(peer.Asn)
			AddToGraph(g, as.Data.Ipv4Upstreams, peer.Asn, deepLevel-1)
		} else {
			return
		}
	}
	if debugEnv {
		fmt.Println(g)
	}
}
