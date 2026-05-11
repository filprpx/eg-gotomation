package models

type NestedManufacturer struct {
	APIBaseFields

	DeviceTypeCount int `json:"devicetype_count"`
}

type manufacturerCommon struct {
	Tags *[]NestedTag `json:"tags"`
}

type Manufacturer struct {
	NestedManufacturer
	manufacturerCommon

	InventoryItemCount int `json:"inventoryitem_count"`
	PlatformCount      int `json:"platform_count"`
}

func (m Manufacturer) MapToWrite() APIResourceWrite {
	return &ManufacturerWrite{
		APIWriteBaseFields: *m.APIBaseFields.MapToWrite(),
		manufacturerCommon: m.manufacturerCommon,
	}
}

type ManufacturerWrite struct {
	APIWriteBaseFields
	manufacturerCommon
}
