package types

type AsnPeer struct {
  Asn int64 `json:"asn"`
  Name string `json:"name"`
  Description string `json:"description"`
  CountryCode string `json:"country_code"`
}

type AsnPeers struct {
  Ipv4Peers []AsnPeer `json:"ipv4_peers"`
  Ipv6Peers []AsnPeer `json:"ipv6_peers"`
}
