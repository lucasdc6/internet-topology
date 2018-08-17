package main

import (
  "fmt"
  "github.com/openrdap/rdap"
  "encoding/json"
  "testing"
)

func TestRdap(t *testing.T) {
  client := &rdap.Client{}
  as, err := client.QueryAutnum("5692")
  if err == nil {
    json, err := json.MarshalIndent(as, "", "  ")
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(string(json))
  }
}
