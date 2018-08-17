package types

type Asn struct {
  Asn int64 `json:"asn"`
  Name string `json:"name"`
  DescriptionShort string `json:"description_short"`
  DescriptionFull []string `json:"description_full"`
  CountryCode string `json:"country_code"`
  Website string `json:"website"`
  EmailContacts []string `json:"email_contacts"`
  AbuseContacts []string `json:"abuse_contacts"`
  LookingGlass string `json:"looking_glass"`
  TrafficEstimation string `json:"traffic_estimation"`
  TrafficRatio string `json:"traffic_ratio"`
  OwnerAddress []string `json:"owner_address"`
  RirAllocation struct {
    RirName string `json:"rir_name"`
    CountryCode string `json:"country_code"`
    DateAllocated string `json:"date_allocated"`
  } `json:"rir_allocation"`
  DateUpdated string `json:"date_updated"`
}