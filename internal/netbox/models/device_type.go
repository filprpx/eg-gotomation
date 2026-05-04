package models

type NestedDeviceType struct {
	APIBaseFields
	Manufacturer NestedManufacturer `json:"manufacturer"`
	Model        string             `json:"model"`
}

type DeviceType struct {
	APIBaseFields
	Manufacturer NestedManufacturer `json:"manufacturer"`
	Model        string             `json:"model"`
	DeviceCount  int                `json:"device_count"`
}

func (d DeviceType) MapToWrite() APIResourceWrite {
	return &DeviceTypeWrite{
		APIWriteBaseFields: *d.APIBaseFields.MapToWrite(),
		Manufacturer:       d.Manufacturer.Id,
	}
}

type DeviceTypeWrite struct {
	APIWriteBaseFields
	Manufacturer int `json:"manufacturer,omitempty"`
}
