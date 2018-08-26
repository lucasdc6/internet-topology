package test

import (
	"fmt"
	"github.com/lucasdc6/internet-topology/peeringdb/api"
	"testing"
)

func TestQueries(t *testing.T) {
	asn, err := api.GetAsn(5692)
	if err != nil {
		panic(err)
	}

	network, err := api.GetNetwork(asn.Data[0].Id)
	if err != nil {
		panic(err)
	}

	ix, err := api.GetIx(network.Data[0].NetixlanSet[0].IxlanId)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n\n", asn.Data)
	fmt.Printf("%+v\n\n", network.Data)
	fmt.Printf("%+v\n\n", ix.Data[0].Org.NetSet)
}
