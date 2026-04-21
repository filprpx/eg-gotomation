package netbox

import (
	"context"
)

type CableAPI struct {
	client *NetboxClient
}

type Cable struct{}

func NewCableAPI(client *NetboxClient) *CableAPI {
	return &CableAPI{
		client: client,
	}
}

func (a *CableAPI) List(ctx context.Context) ([]Cable, error) {
	cables := []Cable{}
	return cables, nil
}
