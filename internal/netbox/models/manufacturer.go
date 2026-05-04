package models

type NestedManufacturer struct {
	ApiBaseFields
}

type Manufacturer struct {
	ApiBaseFields
}

func (m Manufacturer) MapToWrite() ApiResourceWrite {
	return &ManufacturerWrite{
		ApiWriteBaseFields: *m.ApiBaseFields.MapToWrite(),
	}
}

type ManufacturerWrite struct {
	ApiWriteBaseFields
}
