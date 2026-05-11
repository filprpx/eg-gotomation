package models

type NestedPlatform struct {
	APIBaseFields
	DeviceCount         int `json:"device_count"`
	VirtualMachineCount int `json:"virtualmachine_count"`
	Depth               int `json:"_depth"`
}

type platformCommon struct {
	Tags *[]NestedTag `json:"tags"`
}

type Platform struct {
	NestedPlatform
	platformCommon

	Parent         *NestedPlatform       `json:"parent"`
	Manufacturer   *NestedManufacturer   `json:"manufacturer"`
	ConfigTemplate *NestedConfigTemplate `json:"config_template"`
}

func (s Platform) MapToWrite() APIResourceWrite {
	return &PlatformWrite{
		APIWriteBaseFields: *s.APIBaseFields.MapToWrite(),
		platformCommon:     s.platformCommon,
		Comments:           s.Comments,
		Parent:             SafeGetId(s.Parent),
		Manufacturer:       SafeGetId(s.Manufacturer),
		ConfigTemplate:     SafeGetId(s.ConfigTemplate),
	}
}

type PlatformWrite struct {
	APIWriteBaseFields
	platformCommon
	Comments string `json:"comments,omitempty"`

	Parent         int `json:"parent,omitempty"`
	Manufacturer   int `json:"manufacturer,omitempty"`
	ConfigTemplate int `json:"config_template,omitempty"`
}
