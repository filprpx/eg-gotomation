package models

type NestedRack struct {
	APIBaseFields
}

type rackCommon struct {
	FacilityID    string       `json:"facility_id"`
	Serial        *string      `json:"serial"`
	AssetTAG      *string      `json:"asset_tag"`
	UHeight       int          `json:"u_height"`
	StartingUnit  int          `json:"starting_unit"`
	Weight        *float64     `json:"weight"`
	MaxWeight     *int         `json:"max_weight"`
	DescUnits     bool         `json:"desc_units"`
	OuterWidth    *int         `json:"outer_width"`
	OuterHeight   *int         `json:"outer_height"`
	OuterDepth    *int         `json:"outer_depth"`
	MountingDepth *int         `json:"mounting_depth"`
	Tags          *[]NestedTag `json:"tags"`
}

type Rack struct {
	NestedRack
	rackCommon
	Status *Choice `json:"status"`

	Site       NestedSite        `json:"site"`
	Location   *NestedLocation   `json:"location"`
	Tenant     *NestedTenant     `json:"tenant"`
	Role       *NestedRackRole   `json:"role"`
	RackType   *NestedDeviceType `json:"rack_type"`
	FormFactor *Choice           `json:"form_factor"`
	Width      *Choice           `json:"width"`
	WeightUnit *Choice           `json:"weight_unit"`
	OuterUnit  *Choice           `json:"outer_unit"`
	Airflow    *Choice           `json:"airflow"`

	DeviceCount    int `json:"device_count"`
	PowerfeedCount int `json:"powerfeed_count"`
}

func (s Rack) MapToWrite() APIResourceWrite {
	return &RackWrite{
		APIWriteBaseFields: *s.APIBaseFields.MapToWrite(),
		rackCommon:         s.rackCommon,
		Comments:           s.Comments,
		Status:             safeChoiceValue(s.Status),
		Site:               s.Site.GetId(),
		Location:           SafeGetId(s.Location),
		Tenant:             SafeGetId(s.Tenant),
		Role:               SafeGetId(s.Role),
		RackType:           SafeGetId(s.RackType),
		FormFactor:         safeChoiceValue(s.FormFactor),
		Width:              safeChoiceIntValue(s.Width),
		WeightUnit:         safeChoiceValue(s.WeightUnit),
		OuterUnit:          safeChoiceValue(s.OuterUnit),
		Airflow:            safeChoiceValue(s.Airflow),
	}
}

type RackWrite struct {
	APIWriteBaseFields
	rackCommon
	Comments string `json:"comments,omitempty"`
	Status   string `json:"status,omitempty"`

	Site       int    `json:"site,omitempty"`
	Location   int    `json:"location,omitempty"`
	Tenant     int    `json:"tenant,omitempty"`
	Role       int    `json:"role,omitempty"`
	RackType   int    `json:"rack_type,omitempty"`
	FormFactor string `json:"form_factor,omitempty"`
	Width      int    `json:"width,omitempty"`
	WeightUnit string `json:"weight_unit,omitempty"`
	OuterUnit  string `json:"outer_unit,omitempty"`
	Airflow    string `json:"airflow,omitempty"`
}
