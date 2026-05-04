package models

type NestedLocation struct {
	ApiBaseFields
}

type Location struct {
	ApiBaseFields
}

func (d Location) MapToWrite() ApiResourceWrite {
	return &LocationWrite{
		ApiWriteBaseFields: *d.ApiBaseFields.MapToWrite(),
	}
}

type LocationWrite struct {
	ApiWriteBaseFields
}
