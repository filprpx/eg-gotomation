package models

type NestedTenant struct {
	APIBaseFields
}

type Tenant struct {
	APIBaseFields
}

func (t Tenant) MapToWrite() APIResourceWrite {
	return &TenantWrite{
		APIWriteBaseFields: *t.APIBaseFields.MapToWrite(),
	}
}

type TenantWrite struct {
	APIWriteBaseFields
}
