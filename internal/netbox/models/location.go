package models

type NestedLocation struct {
	ApiBaseFields
}

type Location struct {
	ApiBaseFields
}

func (l *Location) MapToWrite() *LocationWrite {
	return &LocationWrite{
		ApiWriteBaseFields: *l.ApiBaseFields.MapToWrite(),
	}
}

type LocationWrite struct {
	ApiWriteBaseFields
}
