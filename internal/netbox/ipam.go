package netbox

import ()

type IpamAPI struct {
	client *NetboxClient
}

func NewIpamAPI(client *NetboxClient) (*IpamAPI, error) {
	return &IpamAPI{
		client: client,
	}, nil
}
