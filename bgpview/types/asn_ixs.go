package types

// AsnIxs - Asn ixs structure - Get as's ixs list
type AsnIxs struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          []struct {
		IxID        int    `json:"ix_id"`
		Name        string `json:"name"`
		NameFull    string `json:"name_full"`
		CountryCode string `json:"country_code"`
		City        string `json:"city"`
		Ipv4Address string `json:"ipv4_address"`
		Ipv6Address string `json:"ipv6_address"`
		Speed       int    `json:"speed"`
	} `json:"data"`
	Metadata Metadata `json:"@meta"`
}

// toJSON - Print AsnIxs with json format
func (asn AsnIxs) toJSON() string {
	return toJSON(asn.Data)
}
