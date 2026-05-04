package models

type NestedDeviceRole struct {
	ApiBaseFields
}

type DeviceRole struct {
	ApiBaseFields
}

func (d DeviceRole) MapToWrite() ApiResourceWrite {
	return &DeviceRoleWrite{
		ApiWriteBaseFields: *d.ApiBaseFields.MapToWrite(),
	}
}

type DeviceRoleWrite struct {
	ApiWriteBaseFields
}
