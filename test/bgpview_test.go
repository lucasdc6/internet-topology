package test

import (
	"fmt"
	"github.com/lucasdc6/internet-topology/bgpview/api"
	"testing"
)

func TestBgpview(t *testing.T) {
	asn := api.GetAsn(5692)
	asnPrefixes := api.GetAsnPrefixes(5692)
	asnPeers := api.GetAsnPeers(5692)
	asnUpstreams := api.GetAsnUpstreams(5692)
	asnDownstreams := api.GetAsnDownstreams(10834)
	asnIxs := api.GetAsnIxs(5692)
	prefix := api.GetPrefix("192.209.63.0", 24)
	ix := api.GetIx(363)

	fmt.Println(asn)
	fmt.Println(asnPrefixes)
	fmt.Println(asnPeers)
	fmt.Println(asnUpstreams)
	fmt.Println(asnDownstreams)
	fmt.Println(asnIxs)
	fmt.Println(prefix)
	fmt.Println(ix)
}
