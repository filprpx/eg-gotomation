package models

type NestedDeviceRole struct {
	APIBaseFields
	Color               string `json:"color"`
	DeviceCount         int    `json:"device_count"`
	VirtualMachineCount int    `json:"virtualmachine_count"`
	Depth               int    `json:"_depth"`
}

type deviceRoleCommon struct {
	VMRole bool         `json:"vm_role"`
	Tags   *[]NestedTag `json:"tags"`
}

type DeviceRole struct {
	NestedDeviceRole
	deviceRoleCommon

	Parent         *NestedDeviceRole     `json:"parent"`
	ConfigTemplate *NestedConfigTemplate `json:"config_template"`
}

func (d DeviceRole) MapToWrite() APIResourceWrite {
	return &DeviceRoleWrite{
		APIWriteBaseFields: *d.APIBaseFields.MapToWrite(),
		deviceRoleCommon:   d.deviceRoleCommon,
		Comments:           d.Comments,
		Color:              d.Color,
		Parent:             SafeGetId(d.Parent),
		ConfigTemplate:     SafeGetId(d.ConfigTemplate),
	}
}

type DeviceRoleWrite struct {
	APIWriteBaseFields
	deviceRoleCommon
	Comments string `json:"comments,omitempty"`
	Color    string `json:"color,omitempty"`

	Parent         int `json:"parent,omitempty"`
	ConfigTemplate int `json:"config_template,omitempty"`
}
