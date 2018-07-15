package types

type NetworkBody struct {
  Meta Meta
  Data []Network
}

type IxBody struct {
  Meta Meta
  Data []Ix
}

type Meta struct {}

type Network struct {
  Id int64
  Org_Id int64
  Org Organization
  Name string
  Aka string
  Website string
  Asn int64
  Looking_Glass string
  Router_Server string
  Irr_As_Set string
  Info_Type string
  Info_Prefixes4 int64
  Info_Prefixes6 int64
  Info_Traffic string
  Info_Ratio string
  Info_Scope string
  Info_Unicast bool
  Info_Multicast bool
  Info_Ipv_6 bool
  Notes string
  Policy_Url string
  Policy_General string
  Policy_Locations string
  Policy_Ratio bool
  Policy_Contracts string
  Netfac_Set []Netfac
  Netixlan_Set []Netixlan
  Poc_Set []Poc
  Created string
  Updated string
  Status string
}

type Ix struct {
  Id int64
  Org_Id int64
  Org Organization
  Name string
  Name_Long string
  City string
  Country string
  Region_Continent string
  media string
  notes string
  Proto_Unicast bool
  Proto_Multicast bool
  Proto_Ipv6 bool
  Website string
  Url_Stats string
  Tech_Email string
  Tech_Phone string
  Policy_Email string
  Policy_Phone string
  Fac_Set []Fac
  Ixlan_Set []Ixlan
  Created string
  Updated string
  Status string
}

type Organization struct {
  Id int64
  Name string
  Website string
  Notes string
  Net_Set []int64
  Fac_Set []int64
  Ix_Set []int64
  Address1 string
  Address2 string
  City string
  Country string
  State string
  Zipcode string
  Created string
  Updated string
  Status string
}

type Fac struct {
  Id int64
  Org_Id int64
  Org_Name string
  Name string
  Website string
  Clli string
  Rencode string
  Npanxx string
  Notes string
  Net_Count int64
  Latitude float64
  Longitude float64
  Created string
  Updated string
  Status string
  Address1 string
  Address2 string
  City string
  Country string
  State string
  Zipcode string
}

type Netfac struct {}

type Netixlan struct {
  Id int64
  Ix_Id int64
  Name string
  Ixlan_Id int64
  Notes string
  Speed int64
  Asn int64
  Ipaddr4 string
  Ipaddr6 string
  Is_Rs_Peer bool
  Created string
  Updated string
  Status string
}

type Poc struct {}

type Ixlan struct {
  Id int64
  Name string
  Descr string
  Mtu int64
  Dot1q_Support bool
  Rs_Asn int64
  Arp_Sponge struct {}
  Created string
  Updated string
  Status string
}
