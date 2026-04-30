package models

type NestedSite struct {
	ApiBaseFields
}

type Site struct {
	ApiBaseFields
}

func (s *Site) MapToWrite() *SiteWrite {
	return &SiteWrite{
		ApiWriteBaseFields: *s.ApiBaseFields.MapToWrite(),
	}
}

type SiteWrite struct {
	ApiWriteBaseFields
}
