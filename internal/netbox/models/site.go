package models

type NestedSite struct {
	ApiBaseFields
}

type Site struct {
	ApiBaseFields
}

func (s Site) MapToWrite() ApiResourceWrite {
	return &SiteWrite{
		ApiWriteBaseFields: *s.ApiBaseFields.MapToWrite(),
	}
}

type SiteWrite struct {
	ApiWriteBaseFields
}
