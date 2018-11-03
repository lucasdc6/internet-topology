// Generate the internet's topology starting from a specific AS.
//
// Examples:
//
// 1) Basic Example:
//  $ internet-topology --asn=3
// 2) Example with deep 3:
//  $ internet-topolofy --asn=3 --deep=3
// 3) Example with IX connections:
//  $ internet-topology --asn=3 --deep=3 --ix
// For full documentation, use the flag --help
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lucasdc6/internet-topology/graph"
	"github.com/lucasdc6/internet-topology/options"
	"github.com/pborman/getopt/v2"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

func main() {
	// Declare options
	optPipeMode := getopt.BoolLong("pipe", 0, "Set pipe mode.")
	optAsn := getopt.IntLong("asn", 0, -1, "Set AS number")
	optIx := getopt.BoolLong("ix", 'i', "Show IX connections")
	optPeers := getopt.BoolLong("peers", 'p', "Show Peers connections")
	optHelp := getopt.BoolLong("help", 'h', "Show this help")
	optDeepLevel := getopt.IntLong("deep", 0, 1, "Set the deep level. Default to full path")
	optOutput := getopt.StringLong("output", 'o', "/tmp/internet-topology", "Set the output file. Default to /tmp/internet-topology[.dot|.data]")

	// Parse
	getopt.Parse()

	// Check args
	if len(os.Args) < 2 || *optHelp {
		getopt.Usage()
		os.Exit(0)
	}

	// Init graph
	g := simple.NewDirectedGraph()

	if *optAsn > 0 {
		options.HandleAsnOption(*optPipeMode, *optAsn, *optDeepLevel, g)
	}

	if *optIx {
		options.HandleIxOption(*optPipeMode, *optAsn, g)
	}

	if *optPeers {
		options.HandlePeersOption(*optPipeMode, *optAsn, g)
	}

	if len(g.Edges()) > 0 {
		if *optPipeMode {
			fmt.Printf("%s", graph.GraphToData(g, *optAsn))
		} else {
			fmt.Printf("Saving files:\n\t%s.dot\n\t%s.data\n", *optOutput, *optOutput)
		}

		str, _ := dot.Marshal(g, "", "", "  ", false)
		ioutil.WriteFile(fmt.Sprintf("%s.dot", *optOutput), str, 0644)
		ioutil.WriteFile(fmt.Sprintf("%s.data", *optOutput), []byte(graph.GraphToData(g, *optAsn)), 0644)
	} else {
		if !*optPipeMode {
			fmt.Println("Nothing to show")
		}
	}
}
