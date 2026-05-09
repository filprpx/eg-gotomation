// Package netbox is used to interface with netbox API
package netbox

import (
	"github.com/filprpx/eg-gotomation/internal/netbox/api"
	"github.com/filprpx/eg-gotomation/internal/netbox/api/dcim"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type Client struct {
	restapi.Client
	Auth

	Raw *api.RawAPI

	Cable        *dcim.CableAPI
	Device       *dcim.DeviceAPI
	DeviceRole   *dcim.DeviceRoleAPI
	DeviceType   *dcim.DeviceTypeAPI
	Manufacturer *dcim.ManufacturerAPI
	Site         *dcim.SiteAPI
}

func (c *Client) PrepareHeaders() {
	c.Header.Add("Authorization", "Token "+c.APIKey)
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
		Client: *restapi.NewClientWithConfig(&cfg.Config),
		Auth:   cfg.Auth,
	}

	client.PrepareHeaders()

	client.Raw = api.NewRawAPI(client)

	client.Cable = dcim.NewCableAPI(client)
	client.Device = dcim.NewDeviceAPI(client)
	client.DeviceRole = dcim.NewDeviceRoleAPI(client)
	client.DeviceType = dcim.NewDeviceTypeAPI(client)
	client.Manufacturer = dcim.NewManufacturerAPI(client)
	client.Site = dcim.NewSiteAPI(client)

	return client, nil
}
