package types

type AsnIxs struct {
  IxId int64 `json:"ix_id"`
  Name string `json:"name"`
  NameFull string `json:"name_full"`
  CountryCode string `json:"country_code"`
  City string `json:"city"`
  Ipv4Address string `json:"ipv4_address"`
  Ipv6Address string `json:"ipv6_address"`
  Speed int64 `json:"speed"`
}
