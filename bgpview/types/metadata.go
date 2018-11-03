package types

import (
	"encoding/json"
)

// Metadata - Metadata from response
type Metadata struct {
	TimeZone      string `json:"time_zone"`
	APIVersion    int    `json:"api_version"`
	ExecutionTime string `json:"execution_time"`
}

func toJSON(data interface{}) string {
	strJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(strJSON)
}
