package types

// IxMember - Ix member, modelate an ix
type IxMember struct {
	Asn         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
	Ipv4Address string `json:"ipv4_address"`
	Ipv6Address string `json:"ipv6_address"`
	Speed       int    `json:"speed"`
}

// Ix - Ix, use IxMember in members array
type Ix struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		Name         string     `json:"name"`
		NameFull     string     `json:"name_full"`
		Website      string     `json:"website"`
		TechEmail    string     `json:"tech_email"`
		TechPhone    string     `json:"tech_phone"`
		PolicyEmail  string     `json:"policy_email,omitempty"`
		PolicyPhone  string     `json:"policy_phone"`
		City         string     `json:"city"`
		CountryCode  string     `json:"country_code"`
		URLStats     string     `json:"url_stats,omitempty"`
		MembersCount int        `json:"members_count"`
		Members      []IxMember `json:"members"`
	} `json:"data"`
	Metadata Metadata `json:"@meta"`
}

// toJSON - Print Ix with json format
func (ix Ix) toJSON() string {
	return toJSON(ix)
}
