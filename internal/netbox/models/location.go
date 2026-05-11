package models

type NestedLocation struct {
	APIBaseFields
	RackCount int `json:"rack_count"`
	Depth     int `json:"_depth"`
}

type locationCommon struct {
	Facility string       `json:"facility"`
	Tags     *[]NestedTag `json:"tags"`
}

type Location struct {
	NestedLocation
	locationCommon
	Status *Choice `json:"status"`

	Site   NestedSite      `json:"site"`
	Parent *NestedLocation `json:"parent"`
	Tenant *NestedTenant   `json:"tenant"`

	DeviceCount int `json:"device_count"`
	PrefixCount int `json:"prefix_count"`
}

func (d Location) MapToWrite() APIResourceWrite {
	return &LocationWrite{
		APIWriteBaseFields: *d.APIBaseFields.MapToWrite(),
		locationCommon:     d.locationCommon,
		Comments:           d.Comments,
		Status:             safeChoiceValue(d.Status),
		Site:               d.Site.GetId(),
		Parent:             SafeGetId(d.Parent),
		Tenant:             SafeGetId(d.Tenant),
	}
}

type LocationWrite struct {
	APIWriteBaseFields
	locationCommon
	Comments string `json:"comments,omitempty"`
	Status   string `json:"status,omitempty"`

	Site   int `json:"site,omitempty"`
	Parent int `json:"parent,omitempty"`
	Tenant int `json:"tenant,omitempty"`
}
