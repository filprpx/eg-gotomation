package models

type NestedDeviceRole struct {
	APIBaseFields
}

type DeviceRole struct {
	APIBaseFields
}

func (d DeviceRole) MapToWrite() APIResourceWrite {
	return &DeviceRoleWrite{
		APIWriteBaseFields: *d.APIBaseFields.MapToWrite(),
	}
}

type DeviceRoleWrite struct {
	APIWriteBaseFields
}
