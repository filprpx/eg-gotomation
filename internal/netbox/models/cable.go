package models

type NestedCable struct {
	ApiBaseFields
}

type Cable struct {
	ApiBaseFields
}

func (c Cable) MapToWrite() ApiResourceWrite {
	return &CableWrite{
		ApiWriteBaseFields: *c.ApiBaseFields.MapToWrite(),
	}
}

type CableWrite struct {
	ApiWriteBaseFields
}
