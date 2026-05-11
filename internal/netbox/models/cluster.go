package models

type NestedCluster struct {
	APIBaseFields
}

type clusterCommon struct {
	ScopeType string       `json:"scope_type"`
	ScopeID   int          `json:"scope_id"`
	Tags      *[]NestedTag `json:"tags"`
}

type Cluster struct {
	NestedCluster
	clusterCommon

	Type   *NestedClusterType  `json:"type"`
	Group  *NestedClusterGroup `json:"group"`
	Status *Choice             `json:"status"`
	Tenant *NestedTenant       `json:"tenant"`

	Scope string `json:"scope"`

	DeviceCount         int `json:"device_count"`
	VirtualMachineCount int `json:"virtualmachine_count"`
	AllocatedVCPUs      int `json:"allocated_vcpus"`
	AllocatedMemory     int `json:"allocated_memory"`
	AllocatedDisk       int `json:"allocated_disk"`
}

func (s Cluster) MapToWrite() APIResourceWrite {
	return &ClusterWrite{
		APIWriteBaseFields: *s.APIBaseFields.MapToWrite(),
		clusterCommon:      s.clusterCommon,
		Comments:           s.Comments,
		Type:               SafeGetId(s.Type),
		Group:              SafeGetId(s.Group),
		Status:             safeChoiceValue(s.Status),
		Tenant:             SafeGetId(s.Tenant),
	}
}

type ClusterWrite struct {
	APIWriteBaseFields
	clusterCommon
	Comments string `json:"comments,omitempty"`

	Type   int    `json:"type,omitempty"`
	Group  int    `json:"group,omitempty"`
	Status string `json:"status,omitempty"`
	Tenant int    `json:"tenant,omitempty"`
}
