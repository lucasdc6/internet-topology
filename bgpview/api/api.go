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

func GetAsn(id int64) types.Response {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d", id)
  response := types.Response{ Data: types.Asn{}}
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnPrefixes(id int64) types.Response {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/prefixes", id)
  response := types.Response{ Data: types.AsnPrefixes{}}
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnPeers(id int64) types.Response {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/peers", id)
  response := types.Response{ Data: types.AsnPrefixes{}}
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnUpstreams(id int64) types.Response {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/upstreams", id)
  response := types.Response{ Data: types.AsnUpstreams{}}
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnDownstreams(id int64) types.Response {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/downstreams", id)
  response := types.Response{ Data: types.AsnDownstreams{}}
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnIxs(id int64) types.Response {
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/ixs", id)
  response := types.Response{ Data: types.AsnIxs{}}
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetPrefix(ip string, cidr int64) types.Response {
  url := fmt.Sprintf("https://api.bgpview.io/prefix/%s/%d", ip, cidr)
  response := types.Response{ Data: types.Prefix{}}
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetIx(id int64) types.Response {
  url := fmt.Sprintf("https://api.bgpview.io/ix/%d", id)
  response := types.Response{ Data: types.Ix{}}
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}
