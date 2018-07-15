package main

import (
  "fmt"
  "testing"
  "github.com/lucasdc6/internet-topology/peeringdb/api"
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

  ix, err := api.GetIx(network.Data[0].Netixlan_Set[0].Ixlan_Id)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%+v\n\n", asn.Data)
  fmt.Printf("%+v\n\n", network.Data)
  fmt.Printf("%+v\n\n", ix.Data[0].Org.Net_Set)
}
