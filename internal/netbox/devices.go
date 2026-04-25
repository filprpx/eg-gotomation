package netbox

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DeviceAPI struct {
	client *NetboxClient
}

// struct that receives the data from the api
type Device struct {
	ApiBaseFields
	DeviceType DeviceType `json:"device_type"`
	DeviceRole DeviceRole `json:"role"`
	Site       Site       `json:"site"`
}

func (d *Device) MapToJsonWrite() (json.RawMessage, error) {
	return json.Marshal(DeviceWrite{
		Device:     *d,
		DeviceType: d.DeviceType.Id,
		DeviceRole: d.DeviceRole.Id,
		Site:       d.Site.Id,
	})
}

// struct that prepares the data to the api
type DeviceWrite struct {
	Device
	DeviceType int `json:"device_type"`
	DeviceRole int `json:"role"`
	Site       int `json:"site"`
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

func (a *DeviceAPI) Get(ctx context.Context, id int) (*Device, error) {
	if id == 0 {
		return nil, fmt.Errorf("id can't be zero")
	}

	resp, err := a.client.get(ctx, fmt.Sprintf("/api/dcim/devices/%d", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, a.client.apiError(resp)
	}

	var result Device
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *DeviceAPI) Create(ctx context.Context, device *Device) (*Device, error) {
	if device == nil {
		return nil, fmt.Errorf("device can't be nil")
	}

	payload, err := device.MapToJsonWrite()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	fmt.Println(string(payload))

	resp, err := a.client.post(ctx,
		"/api/dcim/devices/",
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, a.client.apiError(resp)
	}

	var result Device
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *DeviceAPI) Update(ctx context.Context, device *Device) (bool, error) {
	if device == nil {
		return false, fmt.Errorf("device can't be nil")
	}

	if device.Id == 0 {
		return false, fmt.Errorf("device.Id can't be zero")
	}

	payload, err := device.MapToJsonWrite()
	if err != nil {
		return false, fmt.Errorf("failed to marshal: %w", err)
	}

	fmt.Println(string(payload))

	resp, err := a.client.patch(ctx,
		fmt.Sprintf("/api/dcim/devices/%d/", device.Id),
		bytes.NewReader(payload),
	)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, a.client.apiError(resp)
	}

	return true, nil
}

func (a *DeviceAPI) Delete(ctx context.Context, device *Device) (bool, error) {
	if device == nil {
		return false, fmt.Errorf("device can't be nil")
	}

	if device.Id == 0 {
		return false, fmt.Errorf("device.Id can't be zero")
	}

	resp, err := a.client.delete(ctx,
		fmt.Sprintf("/api/dcim/devices/%d", device.Id),
	)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, a.client.apiError(resp)
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	return true, nil
}
