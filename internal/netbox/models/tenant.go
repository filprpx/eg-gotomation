package models

type NestedTenant struct {
	APIBaseFields
}

type tenantCommon struct {
	Tags *[]NestedTag `json:"tags"`
}

type Tenant struct {
	NestedTenant
	tenantCommon

	Group *NestedTenantGroup `json:"group"`

	CircuitCount        int `json:"circuit_count"`
	DeviceCount         int `json:"device_count"`
	IPAddressCount      int `json:"ipaddress_count"`
	PrefixCount         int `json:"prefix_count"`
	RackCount           int `json:"rack_count"`
	SiteCount           int `json:"site_count"`
	VirtualMachineCount int `json:"virtualmachine_count"`
	VLANCount           int `json:"vlan_count"`
	VRFCount            int `json:"vrf_count"`
	ClusterCount        int `json:"cluster_count"`
}

func (t Tenant) MapToWrite() APIResourceWrite {
	return &TenantWrite{
		APIWriteBaseFields: *t.APIBaseFields.MapToWrite(),
		tenantCommon:       t.tenantCommon,
		Comments:           t.Comments,
		Group:              SafeGetId(t.Group),
	}
}

type TenantWrite struct {
	APIWriteBaseFields
	tenantCommon
	Comments string `json:"comments,omitempty"`

	Group int `json:"group,omitempty"`
}
