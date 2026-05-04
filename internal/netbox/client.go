package netbox

import (
	"github.com/filprpx/eg-gotomation/internal/netbox/api/dcim"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type Client struct {
	restapi.BaseClient

	Cable        *dcim.CableAPI
	Device       *dcim.DeviceAPI
	DeviceRole   *dcim.DeviceRoleAPI
	DeviceType   *dcim.DeviceTypeAPI
	Manufacturer *dcim.ManufacturerAPI
	Site         *dcim.SiteAPI
}

func (c *Client) PrepareHeaders() {
	c.Header.Add("Authorization", "Token "+c.ApiKey)
}

func NewClient() (*Client, error) {
	cfg := DefaultConfig()
	return NewClientWithConfig(cfg)
}

func NewClientWithConfig(cfg *Config) (*Client, error) {
	err := ValidateConfig(cfg)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Config: *cfg,
	}

	client.PrepareHeaders()

	client.Cable = dcim.NewCableAPI(client)
	client.Device = dcim.NewDeviceAPI(client)
	client.DeviceRole = dcim.NewDeviceRoleAPI(client)
	client.DeviceType = dcim.NewDeviceTypeAPI(client)
	client.Manufacturer = dcim.NewManufacturerAPI(client)

	return client, nil
}
