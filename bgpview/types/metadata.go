package types

import (
  "encoding/json"
)

// Metadata
type Metadata struct {
  TimeZone string `json:"time_zone"`
  ApiVersion int `json:"api_version"`
  ExecutionTime string `json:"execution_time"`
}

func toJson(data interface{}) string {
  strJson, err := json.MarshalIndent(data, "", "  ")
  if err != nil {
    panic(err)
  }
  return string(strJson)
}
