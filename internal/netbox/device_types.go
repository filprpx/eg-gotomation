package netbox

import (
	"context"
)

type DeviceTypeAPI struct {
	client *NetboxClient
}

type DeviceType struct{}

func NewDeviceTypeAPI(client *NetboxClient) *DeviceTypeAPI {
	return &DeviceTypeAPI{
		client: client,
	}
}

func (a *DeviceTypeAPI) List(ctx context.Context) ([]DeviceType, error) {
	aDeviceTypes := []DeviceType{}
	return aDeviceTypes, nil
}
