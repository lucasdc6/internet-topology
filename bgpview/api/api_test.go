package api

import (
	"fmt"
	"testing"
)

func TestBgpview(t *testing.T) {
	asn := GetAsn(5692)
	asnPrefixes := GetAsnPrefixes(5692)
	asnPeers := GetAsnPeers(5692)
	asnUpstreams := GetAsnUpstreams(5692)
	asnDownstreams := GetAsnDownstreams(10834)
	asnIxs := GetAsnIxs(5692)
	prefix := GetPrefix("192.209.63.0", 24)
	ix := GetIx(363)

	fmt.Println(asn)
	fmt.Println(asnPrefixes)
	fmt.Println(asnPeers)
	fmt.Println(asnUpstreams)
	fmt.Println(asnDownstreams)
	fmt.Println(asnIxs)
	fmt.Println(prefix)
	fmt.Println(ix)
}
