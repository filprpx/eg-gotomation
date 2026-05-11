package models

type NestedDeviceType struct {
	APIBaseFields
	Manufacturer NestedManufacturer `json:"manufacturer"`
	Model        string             `json:"model"`
	DeviceCount  int                `json:"device_count"`
}

type deviceTypeCommon struct {
	PartNumber             string       `json:"part_number"`
	UHeight                int          `json:"u_height"`
	ExcludeFromUtilization bool         `json:"exclude_from_utilization"`
	IsFullDepth            bool         `json:"is_full_depth"`
	Weight                 *float64     `json:"weight"`
	FrontImage             string       `json:"front_image"`
	RearImage              string       `json:"rear_image"`
	Tags                   *[]NestedTag `json:"tags"`
}

type DeviceType struct {
	NestedDeviceType
	deviceTypeCommon

	DefaultPlatform *NestedPlatform `json:"default_platform"`

	SubdeviceRole *Choice `json:"subdevice_role"`
	Airflow       *Choice `json:"airflow"`
	WeightUnit    *Choice `json:"weight_unit"`

	ConsolePortTemplateCount       int `json:"console_port_template_count"`
	ConsoleServerPortTemplateCount int `json:"console_server_port_template_count"`
	PowerPortTemplateCount         int `json:"power_port_template_count"`
	PowerOutletTemplateCount       int `json:"power_outlet_template_count"`
	InterfaceTemplateCount         int `json:"interface_template_count"`
	FrontPortTemplateCount         int `json:"front_port_template_count"`
	RearPortTemplateCount          int `json:"rear_port_template_count"`
	DeviceBayTemplateCount         int `json:"device_bay_template_count"`
	ModuleBayTemplateCount         int `json:"module_bay_template_count"`
	InventoryItemTemplateCount     int `json:"inventory_item_template_count"`
}

func (d DeviceType) MapToWrite() APIResourceWrite {
	return &DeviceTypeWrite{
		APIWriteBaseFields: *d.APIBaseFields.MapToWrite(),
		deviceTypeCommon:   d.deviceTypeCommon,
		Comments:           d.Comments,
		Manufacturer:       d.Manufacturer.Id,
		DefaultPlatform:    SafeGetId(d.DefaultPlatform),
		SubdeviceRole:      safeChoiceValue(d.SubdeviceRole),
		Airflow:            safeChoiceValue(d.Airflow),
		WeightUnit:         safeChoiceValue(d.WeightUnit),
	}
}

type DeviceTypeWrite struct {
	APIWriteBaseFields
	deviceTypeCommon
	Comments string `json:"comments,omitempty"`

	Manufacturer    int    `json:"manufacturer,omitempty"`
	DefaultPlatform int    `json:"default_platform,omitempty"`
	SubdeviceRole   string `json:"subdevice_role,omitempty"`
	Airflow         string `json:"airflow,omitempty"`
	WeightUnit      string `json:"weight_unit,omitempty"`
}
