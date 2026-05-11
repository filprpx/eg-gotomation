package models

type NestedClusterType struct {
	APIBaseFields
	ClusterCount int `json:"cluster_count"`
}
