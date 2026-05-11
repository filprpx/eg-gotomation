package models

type NestedSite struct {
	APIBaseFields
}

type siteCommon struct {
	Facility        string       `json:"facility"`
	TimeZone        string       `json:"time_zone"`
	PhysicalAddress string       `json:"physical_address"`
	ShippingAddress string       `json:"shipping_address"`
	Latitude        *float64     `json:"latitude"`
	Longitude       *float64     `json:"longitude"`
	Tags            *[]NestedTag `json:"tags"`
}

type Site struct {
	NestedSite
	siteCommon
	Status *Choice `json:"status"`

	Region *NestedRegion    `json:"region"`
	Group  *NestedSiteGroup `json:"group"`
	Tenant *NestedTenant    `json:"tenant"`
	ASNs   *[]NestedASN     `json:"asns"`

	CircuitCount        int `json:"circuit_count"`
	DeviceCount         int `json:"device_count"`
	PrefixCount         int `json:"prefix_count"`
	RackCount           int `json:"rack_count"`
	VirtualMachineCount int `json:"virtualmachine_count"`
	VLANCount           int `json:"vlan_count"`
}

func (s Site) MapToWrite() APIResourceWrite {
	return &SiteWrite{
		APIWriteBaseFields: *s.APIBaseFields.MapToWrite(),
		siteCommon:         s.siteCommon,
		Comments:           s.Comments,
		Status:             safeChoiceValue(s.Status),
		Region:             SafeGetId(s.Region),
		Group:              SafeGetId(s.Group),
		Tenant:             SafeGetId(s.Tenant),
	}
}

type SiteWrite struct {
	APIWriteBaseFields
	siteCommon
	Comments string `json:"comments,omitempty"`
	Status   string `json:"status,omitempty"`

	Region int `json:"region,omitempty"`
	Group  int `json:"group,omitempty"`
	Tenant int `json:"tenant,omitempty"`
}
