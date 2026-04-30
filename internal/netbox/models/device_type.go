package models

type NestedDeviceType struct {
	ApiBaseFields
	Manufacturer NestedManufacturer `json:"manufacturer"`
	Model        string             `json:"model"`
	Slug         string             `json:"slug"`
	Description  string             `json:"description"`
}

type DeviceType struct {
	ApiBaseFields
	Manufacturer NestedManufacturer `json:"manufacturer"`
	Model        string             `json:"model"`
	Slug         string             `json:"slug"`
	Description  string             `json:"description"`
	DeviceCount  int                `json:"device_count"`
}

func (d *DeviceType) MapToWrite() *DeviceTypeWrite {
	return &DeviceTypeWrite{
		ApiWriteBaseFields: *d.ApiBaseFields.MapToWrite(),
		Manufacturer:       d.Manufacturer.Id,
	}
}

type DeviceTypeWrite struct {
	ApiWriteBaseFields
	Manufacturer int `json:"manufacturer,omitempty"`
}
