package types

// NetworkBody
type NetworkBody struct {
  Meta Meta
  Data []Network
}

// IxBody
type IxBody struct {
  Meta Meta
  Data []Ix
}

// Meta - Metadata
type Meta struct {}

// Network - Network response
type Network struct {
  Id int64 `json:"id"`
  OrgId int64 `json:"org_id"`
  Org Organization `json:"org"`
  Name string `json:"name"`
  Aka string `json:"aka"`
  Website string `json:"website"`
  Asn int64 `json:"asn"`
  LookingGlass string `json:"looking_glass"`
  RouterServer string `json:"router_server"`
  IrrAsSet string `json:"irr_as_set"`
  InfoType string `json:"info_type"`
  InfoPrefixes4 int64 `json:"info_prefixes4"`
  InfoPrefixes6 int64 `json:"info_prefixes6"`
  InfoTraffic string `json:"info_traffic"`
  InfoRatio string `json:"info_ratio"`
  InfoScope string `json:"info_scope"`
  InfoUnicast bool `json:"info_unicast"`
  InfoMulticast bool `json:"info_multicast"`
  InfoIpv6 bool `json:"info_ipv_6"`
  Notes string `json:"notes"`
  PolicyURL string `json:"policy_url"`
  PolicyGeneral string `json:"policy_general"`
  PolicyLocations string `json:"policy_locations"`
  PolicyRatio bool `json:"policy_ratio"`
  PolicyContracts string `json:"policy_contracts"`
  NetfacSet []Netfac `json:"netfac_set"`
  NetixlanSet []Netixlan `json:"netixlan_set"`
  PocSet []Poc `json:"poc_set"`
  Created string `json:"created"`
  Updated string `json:"updated"`
  Status string `json:"status"`
}

type Ix struct {
  Id int64 `json:"id"`
  OrgId int64 `json:"org_id"`
  Org Organization `json:"org"`
  Name string `json:"name"`
  NameLong string `json:"name_long"`
  City string `json:"city"`
  Country string `json:"country"`
  RegionContinent string `json:"region_continent"`
  Media string `json:"media"`
  Notes string `json:"notes"`
  ProtoUnicast bool `json:"proto_unicast"`
  ProtoMulticast bool `json:"proto_multicast"`
  ProtoIpv6 bool `json:"proto_ipv6"`
  Website string `json:"website"`
  UrlStats string `json:"url_status"`
  TechEmail string `json:"tech_email"`
  TechPhone string `json:"tech_phone"`
  PolicyEmail string `json:"policy_email"`
  PolicyPhone string `json:"policy_phone"`
  FacSet []Fac `json:"fac_set"`
  IxlanSet []Ixlan `json:"ixlan_set"`
  Created string `json:"created"`
  Updated string `json:"updated"`
  Status string `json:"status"`
}

// Organization - Internet's Organization
type Organization struct {
  Id int64 `json:"id"`
  Name string `json:"name"`
  Website string `json:"website"`
  Notes string `json:"notes"`
  NetSet []int64 `json:"net_set"`
  FacSet []int64 `json:"fac_set"`
  IxSet []int64 `json:"ix_set"`
  Address1 string `json:"address1"`
  Address2 string `json:"address2"`
  City string `json:"city"`
  Country string `json:"country"`
  State string `json:"state"`
  Zipcode string `json:"zipcode"`
  Created string `json:"created"`
  Updated string `json:"updated"`
  Status string `json:"status"`
}

// Fac
type Fac struct {
  Id int64 `json:"id"`
  OrgId int64 `json:"org_id"`
  OrgName string `json:"org_name"`
  Name string `json:"name"`
  Website string `json:"website"`
  Clli string `json:"clli"`
  Rencode string `json:"rencode"`
  Npanxx string `json:"npanxx"`
  Notes string `json:"notes"`
  NetCount int64 `json:"net_count"`
  Latitude float64 `json:"latitude"`
  Longitude float64 `json:"longitude"`
  Created string `json:"created"`
  Updated string `json:"updated"`
  Status string `json:"status"`
  Address1 string `json:"address1"`
  Address2 string `json:"address2"`
  City string `json:"city"`
  Country string `json:"country"`
  State string `json:"state"`
  Zipcode string `json:"zipcode"`
}

// Netfac
type Netfac struct {}

// Netixlan
type Netixlan struct {
  Id int64 `json:"id"`
  IxId int64 `json:"ix_id"`
  Name string `json:"name"`
  IxlanId int64 `json:"ixlan_id"`
  Notes string `json:"notes"`
  Speed int64 `json:"speed"`
  Asn int64 `json:"asn"`
  Ipaddr4 string `json:"ipaddr4"`
  Ipaddr6 string `json:"ipaddr6"`
  IsRsPeer bool `json:"is_rs_peer"`
  Created string `json:"created"`
  Updated string `json:"updated"`
  Status string `json:"status"`
}

// Poc
type Poc struct {}

// Ixlan
type Ixlan struct {
  Id int64 `json:"id"`
  Name string `json:"name"`
  Description string `json:"descr"`
  Mtu int64 `json:"mtu"`
  Dot1qSupport bool `json:"dotlq_support"`
  RsAsn int64 `json:"rs_asn"`
  ArpSponge struct {} `json:"arp_sponge"`
  Created string `json:"created"`
  Updated string `json:"updated"`
  Status string `json:"status"`
}
