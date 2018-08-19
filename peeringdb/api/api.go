package api

import (
  "fmt"
  "github.com/lucasdc6/internet-topology/peeringdb/types"
  "github.com/lucasdc6/internet-topology/utils"
)

// Public api
func GetAsn(id int64) (asn types.NetworkBody, error error) {
  url := fmt.Sprintf("https://peeringdb.com/api/net?asn=%d", id)
  error = utils.GetGeneric(url, &asn)
  return asn, error
}

func GetNetwork(id int64) ( network types.NetworkBody, error error) {
  url := fmt.Sprintf("https://peeringdb.com/api/net/%d", id)
  error = utils.GetGeneric(url, &network)
  return network, error
}

func GetIx(id int64) (ix types.IxBody, error error) {
  url := fmt.Sprintf("https://peeringdb.com/api/ix/%d", id)
  error = utils.GetGeneric(url, &ix)
  return ix, error
}
