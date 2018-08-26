package types

// AsnPrefix - Asn prefix, used by AsnPrefixes
type AsnPrefix struct {
	Prefix      string `json:"prefix"`
	Ip          string `json:"ip"`
	Cidr        int64  `json:"cidr"`
	RoaStatus   string `json:"roa_status"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
	Parent      struct {
		Prefix           string `json:"prefix"`
		Ip               string `json:"ip"`
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
		Ipv4Prefixes []AsnPrefix `json:"ipv4_prefixes"`
		Ipv6Prefixes []AsnPrefix `json:"ipv6_prefixes"`
	} `json:"data"`
	Metadata Metadata `json:"@meta"`
}

// ToJson - Print AsnPrefixes with json format
func (asn AsnPrefixes) ToJson() string {
	return toJson(asn)
}
