package types

type AsnIxs struct {
  Status string `json:"status"`
  StatusMessage string `json:"status_message"`
  Data []struct {
    IxId int `json:"ix_id"`
    Name string `json:"name"`
    NameFull string `json:"name_full"`
    CountryCode string `json:"country_code"`
    City string `json:"city"`
    Ipv4Address string `json:"ipv4_address"`
    Ipv6Address string `json:"ipv6_address"`
    Speed int `json:"speed"`
  } `json:"data"`
  Metadata Metadata `json:"@meta"`
}

func (asn AsnIxs) ToJson() string {
  return toJson(asn.Data)
}
