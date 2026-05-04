package models

type NestedSite struct {
	APIBaseFields
}

type Site struct {
	APIBaseFields
}

func (s Site) MapToWrite() APIResourceWrite {
	return &SiteWrite{
		APIWriteBaseFields: *s.APIBaseFields.MapToWrite(),
	}
}

type SiteWrite struct {
	APIWriteBaseFields
}
