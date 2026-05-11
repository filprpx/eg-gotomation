package models

type NestedRegion struct {
	APIBaseFields
	Depth int `json:"_depth"`
}

type regionCommon struct {
	Tags *[]NestedTag `json:"tags"`
}

type Region struct {
	NestedRegion
	regionCommon

	Parent      *NestedRegion `json:"parent"`
	SiteCount   int           `json:"site_count"`
	PrefixCount int           `json:"prefix_count"`
}

func (r Region) MapToWrite() APIResourceWrite {
	return &RegionWrite{
		APIWriteBaseFields: *r.APIBaseFields.MapToWrite(),
		regionCommon:       r.regionCommon,
		Comments:           r.Comments,
		Parent:             SafeGetId(r.Parent),
	}
}

type RegionWrite struct {
	APIWriteBaseFields
	regionCommon
	Comments string `json:"comments,omitempty"`

	Parent int `json:"parent,omitempty"`
}
