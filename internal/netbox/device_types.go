package netbox

import (
	"context"
)

type DeviceTypesAPI struct {
	client *NetboxClient
}

type DeviceType struct {
	ApiBaseFields
	Model       string `json:"model"`
	Slug        string `json:"slug"`
	DeviceCount int    `json:"device_count"`
}

func NewDeviceTypeAPI(client *NetboxClient) *DeviceTypesAPI {
	return &DeviceTypesAPI{
		client: client,
	}
}

func (a *DeviceTypesAPI) List(ctx context.Context) ([]DeviceType, error) {
	aDeviceTypes := []DeviceType{}
	return aDeviceTypes, nil
}
