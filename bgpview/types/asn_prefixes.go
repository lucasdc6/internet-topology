package types

// AsnPrefix - Asn prefix, used by AsnPrefixes
type AsnPrefix struct {
	Prefix      string      `json:"prefix"`
	Ip          string      `json:"ip"`
	Cidr        int64       `json:"cidr"`
	RoaStatus   string      `json:"roa_status,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	CountryCode string      `json:"country_code,omitempty"`
	RirName     string      `json:"rir_name,omitempty"`
	Parent      []AsnPrefix `json:"parent"`
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
