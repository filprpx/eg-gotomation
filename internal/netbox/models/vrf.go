package models

type NestedVRF struct {
	APIBaseFields
	RD          string `json:"rd"`
	PrefixCount int    `json:"prefix_count"`
}
