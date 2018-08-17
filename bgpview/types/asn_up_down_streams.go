package types

type AsnUpDownstream struct {
  Asn int64 `json:"asn"`
  Name string `json:"name"`
  Description string `json:"description"`
  CountryCode string `json:"country_code"`
  BgpPaths []string `json:"bgp_paths"`
}

type AsnUpstreams struct {
  Ipv4Upstreams []AsnUpDownstream `json:"ipv4_upstreams"`
  Ipv6Upstreams []AsnUpDownstream `json:"ipv6_upstreams"`
  Ipv4Graph struct {} `json:"ipv4_graph"`
  Ipv6Graph struct {} `json:"ipv6_graph"`
  CombinedGraph struct {} `json:"combined_graph"`
}

type AsnDownstreams struct {
  Ipv4Downstreams []AsnUpDownstream `json:"ipv4_downstreams"`
  Ipv6Downstreams []AsnUpDownstream `json:"ipv6_downstreams"`
}
