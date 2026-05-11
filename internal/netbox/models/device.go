package models

type NestedDevice struct {
	APIBaseFields
}

// deviceCommon ONLY contains fields where read type == write type
type deviceCommon struct {
	Status    *Choice      `json:"status"`
	Airflow   *Choice      `json:"airflow"`
	Serial    *string      `json:"serial"`
	AssetTAG  *string      `json:"asset_tag"`
	Position  *float64     `json:"position"`
	Face      *Choice      `json:"face"`
	Latitude  *float64     `json:"latitude"`
	Longitude *float64     `json:"longitude"`
	Tags      *[]NestedTag `json:"tags"`
}

type Device struct {
	APIBaseFields
	deviceCommon

	// Fields that change type in write — declared directly
	DeviceType   NestedDeviceType `json:"device_type"`
	Role         NestedDeviceRole `json:"role"`
	Site         NestedSite       `json:"site"`
	Tenant       *NestedTenant    `json:"tenant"`
	Location     *NestedLocation  `json:"location"`
	Rack         *NestedRack      `json:"rack"`
	ParentDevice *NestedDevice    `json:"parent_device"`
	Platform     *NestedPlatform  `json:"platform"`
	Cluster      *NestedCluster   `json:"cluster"`
	PrimaryIP    *NestedIPAddress `json:"primary_ip"`
	PrimaryIP4   *NestedIPAddress `json:"primary_ip4"`
	PrimaryIP6   *NestedIPAddress `json:"primary_ip6"`

	// Read-only count fields
	ConsolePortCount       int `json:"console_port_count"`
	ConsoleServerPortCount int `json:"console_server_port_count"`
	PowerPortCount         int `json:"power_port_count"`
	InterfaceCount         int `json:"interface_count"`
	FrontPortCount         int `json:"front_port_count"`
	DeviceBayCount         int `json:"device_bay_count"`
	ModuleBayCount         int `json:"module_bay_count"`
	InventoryItemCount     int `json:"inventory_item_count"`
}

func (d Device) MapToWrite() APIResourceWrite {
	return &DeviceWrite{
		APIWriteBaseFields: *d.APIBaseFields.MapToWrite(),
		deviceCommon:       d.deviceCommon,

		DeviceType:   d.DeviceType.GetId(),
		Role:         d.Role.GetId(),
		Site:         d.Site.GetId(),
		Tenant:       SafeGetId(d.Tenant),
		Location:     SafeGetId(d.Location),
		Rack:         SafeGetId(d.Rack),
		ParentDevice: SafeGetId(d.ParentDevice),
		Platform:     SafeGetId(d.Platform),
		Cluster:      SafeGetId(d.Cluster),
		PrimaryIP:    SafeGetId(d.PrimaryIP),
		PrimaryIP4:   SafeGetId(d.PrimaryIP4),
		PrimaryIP6:   SafeGetId(d.PrimaryIP6),
	}
}

type DeviceWrite struct {
	APIWriteBaseFields
	deviceCommon

	// Same fields as Device, but with write types
	DeviceType   int `json:"device_type,omitempty"`
	Role         int `json:"role,omitempty"`
	Site         int `json:"site,omitempty"`
	Tenant       int `json:"tenant,omitempty"`
	Location     int `json:"location,omitempty"`
	Rack         int `json:"rack,omitempty"`
	ParentDevice int `json:"parent_device,omitempty"`
	Platform     int `json:"platform,omitempty"`
	Cluster      int `json:"cluster,omitempty"`
	PrimaryIP    int `json:"primary_ip,omitempty"`
	PrimaryIP4   int `json:"primary_ip4,omitempty"`
	PrimaryIP6   int `json:"primary_ip6,omitempty"`
}
