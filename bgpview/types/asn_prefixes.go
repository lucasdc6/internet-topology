package types

type AsnPrefix struct {
  Prefix string `json:"prefix"`
  Ip string `json:"ip"`
  Cidr int64 `json:"cidr"`
  RoaStatus string `json:"roa_status,omitempty"`
  Name string `json:"name,omitempty"`
  Description string `json:"description,omitempty"`
  CountryCode string `json:"country_code,omitempty"`
  RirName string `json:"rir_name,omitempty"`
  Parent []AsnPrefix `json:"parent"`
}

type AsnPrefixes struct {
  Ipv4Prefixes []AsnPrefix `json:"ipv4_prefixes"`
  Ipv6Prefixes []AsnPrefix `json:"ipv6_prefixes"`
}
