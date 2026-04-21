package netbox

import (
	"context"
	"encoding/json"
	"net/http"
)

type DeviceAPI struct {
	client *NetboxClient
	dcim   *DcimAPI

	path string
}

type Device struct {
	ApiBaseFields
	Name       string     `json:"name"`
	DeviceType DeviceType `json:"device_type"`
}

func NewDeviceAPI(client *NetboxClient) *DeviceAPI {
	return &DeviceAPI{
		client: client,
	}
}

func (a *DeviceAPI) List(ctx context.Context) ([]Device, error) {
	resp, err := a.client.get(ctx, "/api/dcim/devices")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, a.client.apiError(resp)
	}

	var result ApiListResponse[Device]
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}
