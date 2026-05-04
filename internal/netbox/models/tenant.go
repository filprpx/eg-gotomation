package models

type NestedTenant struct {
	ApiBaseFields
}

type Tenant struct {
	ApiBaseFields
}

func (t Tenant) MapToWrite() ApiResourceWrite {
	return &TenantWrite{
		ApiWriteBaseFields: *t.ApiBaseFields.MapToWrite(),
	}
}

type TenantWrite struct {
	ApiWriteBaseFields
}
