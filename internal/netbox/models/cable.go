package models

type NestedCable struct {
	APIBaseFields
}

type Cable struct {
	APIBaseFields
}

func (c Cable) MapToWrite() APIResourceWrite {
	return &CableWrite{
		APIWriteBaseFields: *c.APIBaseFields.MapToWrite(),
	}
}

type CableWrite struct {
	APIWriteBaseFields
}
