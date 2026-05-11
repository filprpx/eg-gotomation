package models

type NestedClusterGroup struct {
	APIBaseFields
	ClusterCount int `json:"cluster_count"`
}
