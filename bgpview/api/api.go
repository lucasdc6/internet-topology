package api

import (
  "os"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "github.com/lucasdc6/internet-topology/bgpview/types"
  "github.com/lucasdc6/internet-topology/environment"
)

var debug = environment.GetDebugFor("bgpview") || environment.GetDebug()

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

  if debug {
    fmt.Printf("Query for %s\nbody: %v", url, string(body))
  }

  err = json.Unmarshal([]byte(body), &res)

  if err != nil {
    error = err
  }

  return error
}

func GetAsn(id int) (response types.Asn) {
  if debug {
    fmt.Printf("GetAsn query with: %d\n", id)
  }
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnPrefixes(id int) (response types.AsnPrefixes) {
  if debug {
    fmt.Printf("GetAsnPrefixes query with: %d\n", id)
  }
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/prefixes", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnPeers(id int) (response types.AsnPeers) {
  if debug {
    fmt.Printf("GetAsnPeers query with: %d\n", id)
  }
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/peers", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnUpstreams(id int) (response types.AsnUpstreams) {
  if debug {
    fmt.Printf("GetAsnUpstreams query with: %d\n", id)
  }
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/upstreams", id)
  err := getGeneric(url, &response)
  if err != nil {
    fmt.Printf("Error quering\n")
    os.Exit(2)
  }
  return response
}

func GetAsnDownstreams(id int) (response types.AsnDownstreams) {
  if debug {
    fmt.Printf("GetAsnDownstreams query with: %d\n", id)
  }
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/downstreams", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetAsnIxs(id int) (response types.AsnIxs) {
  if debug {
    fmt.Printf("GetAsnIxs query with: %d\n", id)
  }
  url := fmt.Sprintf("https://api.bgpview.io/asn/%d/ixs", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetPrefix(ip string, cidr int) (response types.Prefix) {
  if debug {
    fmt.Printf("GetPrefix query with: %s and %d\n", ip, cidr)
  }
  url := fmt.Sprintf("https://api.bgpview.io/prefix/%s/%d", ip, cidr)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}

func GetIx(id int) (response types.Ix) {
  if debug {
    fmt.Printf("GetIx query with: %d\n", id)
  }
  url := fmt.Sprintf("https://api.bgpview.io/ix/%d", id)
  err := getGeneric(url, &response)
  if err != nil {
    panic(err)
  }
  return response
}
