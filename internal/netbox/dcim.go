package netbox

import ()

type DcimAPI struct {
	client *NetboxClient

	Cable      *CableAPI
	Device     *DeviceAPI
	DeviceType *DeviceTypeAPI
}

func NewDcimAPI(client *NetboxClient) (*DcimAPI, error) {
	return &DcimAPI{
		client: client,

		Cable:      NewCableAPI(client),
		Device:     NewDeviceAPI(client),
		DeviceType: NewDeviceTypeAPI(client),
	}, nil
}
