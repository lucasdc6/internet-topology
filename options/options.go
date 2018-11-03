package options

import (
	"fmt"

	"github.com/lucasdc6/internet-topology/bgpview/api"
	"github.com/lucasdc6/internet-topology/graph"
	"gonum.org/v1/gonum/graph/simple"
)

// HandleAsnOption -
func HandleAsnOption(pipeMode bool, asn, deepLevel int, g *simple.DirectedGraph) {
	if !pipeMode {
		fmt.Printf("Quering for ASN: %d\n", asn)
	}
	as := api.GetAsnUpstreams(asn)
	graph.AddToGraph(g, as.Data.Ipv4Upstreams, asn, deepLevel)
}

// HandleIxOption -
func HandleIxOption(pipeMode bool, asn int, g *simple.DirectedGraph) {
	if !pipeMode {
		fmt.Printf("Quering for ASN IX: %d\n", asn)
	}
	asIxs := api.GetAsnIxs(asn)
	for _, asIx := range asIxs.Data {
		if !pipeMode {
			fmt.Printf("Quering for IX %s (%d)\n", asIx.Name, asIx.IxID)
		}
		ix := api.GetIx(asIx.IxID)
		for _, member := range ix.Data.Members {
			if member.Asn != asn {
				g.SetEdge(simple.Edge{F: simple.Node(asn), T: simple.Node(member.Asn)})
			}
		}
	}
}

// HandlePeersOption -
func HandlePeersOption(pipeMode bool, asn int, g *simple.DirectedGraph) {
	if !pipeMode {
		fmt.Printf("Quering for ASN Peers: %d\n", asn)
	}
	asPeers := api.GetAsnPeers(asn)
	for _, asPeer := range asPeers.Data.Ipv4Peers {
		g.SetEdge(simple.Edge{F: simple.Node(asn), T: simple.Node(asPeer.Asn)})
	}
}
