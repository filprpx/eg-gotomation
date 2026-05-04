package models

type NestedLocation struct {
	APIBaseFields
}

type Location struct {
	APIBaseFields
}

func (d Location) MapToWrite() APIResourceWrite {
	return &LocationWrite{
		APIWriteBaseFields: *d.APIBaseFields.MapToWrite(),
	}
}

type LocationWrite struct {
	APIWriteBaseFields
}
