package api

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/lucasdc6/internet-topology/peeringdb/types"
)

// Private function
// Get from url and put data in res
// Return only error
func getGeneric(url string, res interface{}) (error error) {
  resp, err := http.Get(url)
  if err != nil {
    error = err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    error = err
  }

  err = json.Unmarshal([]byte(body), &res)

  if err != nil {
    error = err
  }

  return error
}


// Public api
func GetAsn(id int64) (asn types.NetworkBody, error error) {
  url := fmt.Sprintf("https://peeringdb.com/api/net?asn=%d", id)
  error = getGeneric(url, &asn)
  return asn, error
}

func GetNetwork(id int64) ( network types.NetworkBody, error error) {
  url := fmt.Sprintf("https://peeringdb.com/api/net/%d", id)
  error = getGeneric(url, &network)
  return network, error
}

func GetIx(id int64) (ix types.IxBody, error error) {
  url := fmt.Sprintf("https://peeringdb.com/api/ix/%d", id)
  error = getGeneric(url, &ix)
  return ix, error
}
