package models

type NestedManufacturer struct {
	APIBaseFields
}

type Manufacturer struct {
	APIBaseFields
}

func (m Manufacturer) MapToWrite() APIResourceWrite {
	return &ManufacturerWrite{
		APIWriteBaseFields: *m.APIBaseFields.MapToWrite(),
	}
}

type ManufacturerWrite struct {
	APIWriteBaseFields
}
