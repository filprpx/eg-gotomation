package models

type NestedSiteGroup struct {
	APIBaseFields
	Depth int `json:"_depth"`
}

type siteGroupCommon struct {
	Tags *[]NestedTag `json:"tags"`
}

type SiteGroup struct {
	NestedSiteGroup
	siteGroupCommon

	Parent      *NestedSiteGroup `json:"parent"`
	SiteCount   int              `json:"site_count"`
	PrefixCount int              `json:"prefix_count"`
}

func (s SiteGroup) MapToWrite() APIResourceWrite {
	return &SiteGroupWrite{
		APIWriteBaseFields: *s.APIBaseFields.MapToWrite(),
		siteGroupCommon:    s.siteGroupCommon,
		Comments:           s.Comments,
		Parent:             SafeGetId(s.Parent),
	}
}

type SiteGroupWrite struct {
	APIWriteBaseFields
	siteGroupCommon
	Comments string `json:"comments,omitempty"`

	Parent int `json:"parent,omitempty"`
}
