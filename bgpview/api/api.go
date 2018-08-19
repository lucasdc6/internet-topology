package api

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/lucasdc6/internet-topology/bgpview/types"
)

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

func GetAsn(id int64) (response types.Asn) {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnPrefixes(id int64) (response types.AsnPrefixes) {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/prefixes", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnPeers(id int64) (response types.AsnPeers) {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/peers", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnUpstreams(id int64) (response types.AsnUpstreams) {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/upstreams", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnDownstreams(id int64) (response types.AsnDownstreams) {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/downstreams", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnIxs(id int64) (response types.AsnIxs) {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/ixs", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetPrefix(ip string, cidr int64) (response types.Prefix) {
  url := fmt.Sprintf("https://api.bgpview.io/prefix/%s/%d", ip, cidr)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetIx(id int64) (response types.Ix) {
  url := fmt.Sprintf("https://api.bgpview.io/ix/%d", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}
