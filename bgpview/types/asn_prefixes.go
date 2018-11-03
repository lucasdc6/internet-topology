package types

// AsnPrefix - Asn prefix, used by AsnPrefixes
type AsnPrefix struct {
	Prefix      string `json:"prefix"`
	IP          string `json:"ip"`
	Cidr        int64  `json:"cidr"`
	RoaStatus   string `json:"roa_status"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
	Parent      struct {
		Prefix           string `json:"prefix"`
		IP               string `json:"ip"`
		Cidr             int64  `json:"cidr"`
		RirName          string `json:"rir_name"`
		AllocationStatus string `json:"allocation_status"`
	} `json:"parent,omitempty"`
}

// AsnPrefixes - Asn Prefixes, use AsnPrefix - Get as's prefixes
type AsnPrefixes struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		IPv4Prefixes []AsnPrefix `json:"ipv4_prefixes"`
		IPv6Prefixes []AsnPrefix `json:"ipv6_prefixes"`
	} `json:"data"`
	Metadata Metadata `json:"@meta"`
}

// toJSON - Print AsnPrefixes with json format
func (asn AsnPrefixes) toJSON() string {
	return toJSON(asn)
}
