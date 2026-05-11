package models

type NestedTenantGroup struct {
	APIBaseFields
	TenantCount int `json:"tenant_count"`
	Depth       int `json:"_depth"`
}
