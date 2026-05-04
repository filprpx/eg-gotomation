package models

type NestedDeviceType struct {
	ApiBaseFields
	Manufacturer NestedManufacturer `json:"manufacturer"`
	Model        string             `json:"model"`
}

type DeviceType struct {
	ApiBaseFields
	Manufacturer NestedManufacturer `json:"manufacturer"`
	Model        string             `json:"model"`
	DeviceCount  int                `json:"device_count"`
}

func (d DeviceType) MapToWrite() ApiResourceWrite {
	return &DeviceTypeWrite{
		ApiWriteBaseFields: *d.ApiBaseFields.MapToWrite(),
		Manufacturer:       d.Manufacturer.Id,
	}
}

type DeviceTypeWrite struct {
	ApiWriteBaseFields
	Manufacturer int `json:"manufacturer,omitempty"`
}
