package api

import (
	"fmt"
	"github.com/lucasdc6/internet-topology/bgpview/types"
	"github.com/lucasdc6/internet-topology/environment"
	"github.com/lucasdc6/internet-topology/utils"
	"os"
)

var debug = environment.GetDebugFor("bgpview") || environment.GetDebug()

// GetAsn - GetAsn from as number (id)
func GetAsn(id int) (response types.Asn) {
	if debug {
		fmt.Printf("GetAsn query with: %d\n", id)
	}
	url := fmt.Sprintf("https://api.bgpview.io/asn/%d", id)
	err := utils.GetGeneric(url, &response)
	if err != nil {
		panic(err)
	}
	return response
}

// GetAsnPrefixes - Get asn prefixes from as number (id)
func GetAsnPrefixes(id int) (response types.AsnPrefixes) {
	if debug {
		fmt.Printf("GetAsnPrefixes query with: %d\n", id)
	}
	url := fmt.Sprintf("https://api.bgpview.io/asn/%d/prefixes", id)
	err := utils.GetGeneric(url, &response)
	if err != nil {
		panic(err)
	}
	return response
}

// GetAsnPeers - Get as peers from as number (id)
func GetAsnPeers(id int) (response types.AsnPeers) {
	if debug {
		fmt.Printf("GetAsnPeers query with: %d\n", id)
	}
	url := fmt.Sprintf("https://api.bgpview.io/asn/%d/peers", id)
	err := utils.GetGeneric(url, &response)
	if err != nil {
		panic(err)
	}
	return response
}

// GetAsnUpstreams - Get as upstreams from as number (id)
func GetAsnUpstreams(id int) (response types.AsnUpstreams) {
	if debug {
		fmt.Printf("GetAsnUpstreams query with: %d\n", id)
	}
	url := fmt.Sprintf("https://api.bgpview.io/asn/%d/upstreams", id)
	err := utils.GetGeneric(url, &response)
	if err != nil {
    fmt.Println(err)
		fmt.Printf("Error quering\n")
		os.Exit(2)
	}
	return response
}

// GetAsnDownstreams - Get as downstreams from as number (id)
func GetAsnDownstreams(id int) (response types.AsnDownstreams) {
	if debug {
		fmt.Printf("GetAsnDownstreams query with: %d\n", id)
	}
	url := fmt.Sprintf("https://api.bgpview.io/asn/%d/downstreams", id)
	err := utils.GetGeneric(url, &response)
	if err != nil {
		panic(err)
	}
	return response
}

// GetAsnIxs - Get as's ix from as number (id)
func GetAsnIxs(id int) (response types.AsnIxs) {
	if debug {
		fmt.Printf("GetAsnIxs query with: %d\n", id)
	}
	url := fmt.Sprintf("https://api.bgpview.io/asn/%d/ixs", id)
	err := utils.GetGeneric(url, &response)
	if err != nil {
		panic(err)
	}
	return response
}

// GetPrefix - Get internet prefix from ip and cidr
func GetPrefix(ip string, cidr int) (response types.Prefix) {
	if debug {
		fmt.Printf("GetPrefix query with: %s and %d\n", ip, cidr)
	}
	url := fmt.Sprintf("https://api.bgpview.io/prefix/%s/%d", ip, cidr)
	err := utils.GetGeneric(url, &response)
	if err != nil {
		panic(err)
	}
	return response
}

// GetIx - Get internet ix from ix id
func GetIx(id int) (response types.Ix) {
	if debug {
		fmt.Printf("GetIx query with: %d\n", id)
	}
	url := fmt.Sprintf("https://api.bgpview.io/ix/%d", id)
	err := utils.GetGeneric(url, &response)
	if err != nil {
		panic(err)
	}
	return response
}
