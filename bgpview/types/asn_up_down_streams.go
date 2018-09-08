package types

// AsnUpDownstream - Asn up or downstream (generic)
type AsnUpDownstream struct {
	Asn         int      `json:"asn"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CountryCode string   `json:"country_code"`
	BgpPaths    []string `json:"bgp_paths"`
}

// AsnUpstreams - Asn upstreams, use AsnUpDownstream
type AsnUpstreams struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		Ipv4Upstreams []AsnUpDownstream `json:"ipv4_upstreams"`
		Ipv6Upstreams []AsnUpDownstream `json:"ipv6_upstreams"`
		Ipv4Graph     string            `json:"ipv4_graph"`
		Ipv6Graph     string            `json:"ipv6_graph"`
		CombinedGraph string            `json:"combined_graph"`
	} `json:"data"`
	Metadata Metadata `json:"@meta"`
}

// AsnDownstreams - Asn downstreams, use AsnUpDownstream
type AsnDownstreams struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		Ipv4Downstreams []AsnUpDownstream `json:"ipv4_downstreams"`
		Ipv6Downstreams []AsnUpDownstream `json:"ipv6_downstreams"`
	} `json:"data"`
	Metadata Metadata `json:"@meta"`
}

// ToJson - Print AsnUpstreams with json format
func (asn AsnUpstreams) ToJson() string {
	return toJson(asn)
}

// ToJson - Print AsnDownstreams with json format
func (asn AsnDownstreams) ToJson() string {
	return toJson(asn)
}
