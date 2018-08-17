package types

type Response struct {
  Status string `json:"status"`
  StatusMessage string `json:"status_message"`
  Data interface{} `json:"data"`
  Metainformation struct {
    TimeZone string `json:"time_zone"`
    ApiVersion int64 `json:"api_version"`
    ExecutionTime string `json:"execution_time"`
  } `json:"@meta"`
}
