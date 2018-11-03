package types

// AsnPeer - Asn peer structure
type AsnPeer struct {
	Asn         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}

// AsnPeers - Asn peers structure, use AsnPeer
type AsnPeers struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		Ipv4Peers []AsnPeer `json:"ipv4_peers"`
		Ipv6Peers []AsnPeer `json:"ipv6_peers"`
	} `json:"data"`
	Metadata Metadata `json:"@meta"`
}

// toJSON - Print AsnPeers with json format
func (asn AsnPeers) toJSON() string {
	return toJSON(asn)
}
