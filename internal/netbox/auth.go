package netbox

import (
	"fmt"
	"os"
)

const ENV_API_KEY = "NETBOX_API_KEY"

type NetboxAuth struct {
	APIKey string
}

func NewAuth() (*NetboxAuth, error) {
	APIKey := os.Getenv(ENV_API_KEY)
	if APIKey == "" {
		return nil, fmt.Errorf("%s not configured", ENV_API_KEY)
	}

	auth := NetboxAuth{
		APIKey: APIKey,
	}

	return &auth, nil
}
