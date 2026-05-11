package models

type NestedRackRole struct {
	APIBaseFields
	Color string `json:"color"`
}

type rackRoleCommon struct {
	Tags *[]NestedTag `json:"tags"`
}

type RackRole struct {
	NestedRackRole
	rackRoleCommon

	RackCount int `json:"rack_count"`
}

func (r RackRole) MapToWrite() APIResourceWrite {
	return &RackRoleWrite{
		APIWriteBaseFields: *r.APIBaseFields.MapToWrite(),
		rackRoleCommon:     r.rackRoleCommon,
		Color:              r.Color,
	}
}

type RackRoleWrite struct {
	APIWriteBaseFields
	rackRoleCommon
	Color string `json:"color,omitempty"`
}
