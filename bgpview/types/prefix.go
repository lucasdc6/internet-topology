package types

// Prefix Upstream
type PrefixUpstream struct {
  Asn int `json:"asn"`
  Name string `json:"name"`
  Description string `json:"description"`
  CountryCode string `json:"country_code"`
}

// Prefix of internet, us PrefixUpstream
type Prefix struct {
  Status string `json:"status"`
  StatusMessage string `json:"status_message"`
  Data struct {
    Prefix string `json:"prefix"`
    Ip string `json:"ip"`
    Cidr int `json:"cidr"`
    Asns struct {
      Asn int `json:"asn"`
      Name string `json:"name"`
      Description string `json:"description"`
      CountryCode string `json:"country_code"`
      PrefixUpstreams []PrefixUpstream `json:"prefix_upstreams"`
    } `json:"asns"`
    Name string `json:"name"`
    DescriptionShort string `json:"description_short"`
    DescriptionLong []string `json:"description_long"`
    EmailContacts []string `json:"email_contacts"`
    AbuseContacts []string `json:"abuse_contacts"`
    OwnerAddress []string `json:"owner_address"`
    CountryCodes struct {
      WhoisCountryCode string `json:"whois_country_code"`
      RirAllocationCountryCode string `json:"rir_allocation_country_code"`
      MaxmindCountryCode string `json:"maxmind_country_code"`
    } `json:"country_codes"`
    RirAllocation struct {
      RirName string `json:"rir_name"`
      CountryCode string `json:"country_code"`
      Ip string `json:"ip"`
      Cidr int `json:"cidr"`
      Prefix string `json:"prefix"`
      DateAllocated string `json:"date_allocated"`
      AllocationStatus string `json:"allocation_status"`
    } `json:"rir_allocation"`
    IanaAssignment struct {
      AssignmentStatus string `json:"assignment_status"`
      Description string `json:"description"`
      WhoisServer string `json:"whois_server"`
      DateAssigned struct {} `json:"date_assigned"`
    } `json:"iana_assignment"`
    Maxmind struct {
      CountryCode string `json:"country_code"`
      City string `json:"city"`
    } `json:"maxmind"`
    RelatedPrefixes []struct{} `json:"related_prefixes"`
    DateUpdated string `json:"date_updated"`
  } `json:"data"`
  Metadata Metadata `json:"@meta"`
}

func (prefix Prefix) ToJson() string {
  return toJson(prefix)
}

