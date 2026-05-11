package models

type NestedTag struct {
	APIBaseFields
}

type Tag struct {
	APIBaseFields
	Color string `json:"color"`
}

func (s Tag) MapToWrite() APIResourceWrite {
	return &TagWrite{
		APIWriteBaseFields: *s.APIBaseFields.MapToWrite(),
	}
}

type TagWrite struct {
	APIWriteBaseFields
	Color string `json:"color"`
}
