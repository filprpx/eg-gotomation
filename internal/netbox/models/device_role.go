package models

type NestedDeviceRole struct {
	ApiBaseFields
}

type DeviceRole struct {
	ApiBaseFields
}

func (d *DeviceRole) MapToWrite() *DeviceRoleWrite {
	return &DeviceRoleWrite{
		ApiWriteBaseFields: *d.ApiBaseFields.MapToWrite(),
	}
}

type DeviceRoleWrite struct {
	ApiWriteBaseFields
}
