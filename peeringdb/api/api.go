package api

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/lucasdc6/internet-topology/peeringdb/types"
)

func GetAsn(id int64) (asn types.NetworkBody, error error) {
  url := fmt.Sprintf("https://peeringdb.com/api/net?asn=%d", id)
  resp, err := http.Get(url)
  if err != nil {
    error = err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    error = err
  }

  err = json.Unmarshal([]byte(body), &asn)

  if err != nil {
    error = err
  }

  return asn, error
}

func GetNetwork(id int64) ( network types.NetworkBody, error error) {
  url := fmt.Sprintf("https://peeringdb.com/api/net/%d", id)
  resp, err := http.Get(url)
  if err != nil {
    error = err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    error = err
  }

  err = json.Unmarshal([]byte(body), &network)

  if err != nil {
    error = err
  }

  return network, error
}

func GetIx(id int64) (ix types.IxBody, error error) {
  url := fmt.Sprintf("https://peeringdb.com/api/ix/%d", id)
  resp, err := http.Get(url)
  if err != nil {
    error = err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    error = err
  }

  err = json.Unmarshal([]byte(body), &ix)

  if err != nil {
    error = err
  }

  return ix, error
}
