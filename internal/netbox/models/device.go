package models

type NestedDevice struct {
	ApiBaseFields
}

type Device struct {
	ApiBaseFields
	DeviceType NestedDeviceType `json:"device_type"`
	Role       NestedDeviceRole `json:"role"`
	Site       NestedSite       `json:"site"`
	Tenant     NestedTenant     `json:"tenant"`
	Location   NestedLocation   `json:"location"`
	Status     Choice           `json:"status"`
	Airflow    Choice           `json:"airflow"`
}

func (d Device) MapToWrite() ApiResourceWrite {
	return &DeviceWrite{
		ApiWriteBaseFields: *d.ApiBaseFields.MapToWrite(),
		DeviceType:         d.DeviceType.Id,
		Role:               d.Role.Id,
		Site:               d.Site.Id,
		Tenant:             d.Tenant.Id,
		Location:           d.Location.Id,
	}
}

type DeviceWrite struct {
	ApiWriteBaseFields
	DeviceType int `json:"device_type,omitempty"`
	Role       int `json:"role,omitempty"`
	Site       int `json:"site,omitempty"`
	Tenant     int `json:"tenant,omitempty"`
	Location   int `json:"location,omitempty"`
}
