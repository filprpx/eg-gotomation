package models

type NestedASN struct {
	APIBaseFields
	ASN           int           `json:"asn"`
	RIR           *NestedRIR    `json:"rir"`
	Tenant        *NestedTenant `json:"tenant"`
	Tags          *[]NestedTag  `json:"tags"`
	SiteCount     int           `json:"site_count"`
	ProviderCount int           `json:"provider_count"`
}
