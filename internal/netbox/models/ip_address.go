package models

type NestedIPAddress struct {
	APIBaseFields
	Family  any    `json:"family"`
	Address string `json:"address"`
}

type ipAddressCommon struct {
	Address            string       `json:"address"`
	AssignedObjectType string       `json:"assigned_object_type"`
	AssignedObjectID   uint64       `json:"assigned_object_id"`
	DNSName            string       `json:"dns_name"`
	Tags               *[]NestedTag `json:"tags"`
}

type IPAddress struct {
	APIBaseFields
	ipAddressCommon

	Family Choice `json:"family"`

	VRF    *NestedVRF    `json:"vrf"`
	Tenant *NestedTenant `json:"tenant"`

	Status *Choice `json:"status"`
	Role   *Choice `json:"role"`

	AssignedObject string             `json:"assigned_object"`
	NATInside      *NestedIPAddress   `json:"nat_inside"`
	NATOutside     *[]NestedIPAddress `json:"nat_outside"`
}

func (s IPAddress) MapToWrite() APIResourceWrite {
	return &IPAddressWrite{
		APIWriteBaseFields: *s.APIBaseFields.MapToWrite(),
		ipAddressCommon:    s.ipAddressCommon,
		Comments:           s.Comments,
		VRF:                SafeGetId(s.VRF),
		Tenant:             SafeGetId(s.Tenant),
		Status:             safeChoiceValue(s.Status),
		Role:               safeChoiceValue(s.Role),
		NATInside:          SafeGetId(s.NATInside),
	}
}

type IPAddressWrite struct {
	APIWriteBaseFields
	ipAddressCommon
	Comments string `json:"comments,omitempty"`

	VRF       int    `json:"vrf,omitempty"`
	Tenant    int    `json:"tenant,omitempty"`
	Status    string `json:"status,omitempty"`
	Role      string `json:"role,omitempty"`
	NATInside int    `json:"nat_inside,omitempty"`
}
