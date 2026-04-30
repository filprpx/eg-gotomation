package netbox

import (
	"fmt"
	"os"

	"github.com/filprpx/eg-gotomation/internal/netbox/api/dcim"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

const ENV_API_KEY = "NETBOX_API_KEY"

type Client struct {
	restapi.BaseClient

	Device dcim.DeviceAPI
}

func NewClient(baseUrl string) (*Client, error) {
	client := Client{
		BaseClient: *restapi.NewBaseClient(baseUrl),
	}

	err := client.Auth()
	if err != nil {
		return nil, fmt.Errorf("NewClient: %w", err)
	}

	return &client, nil
}

func (c *Client) Auth() error {
	apiKey := os.Getenv(ENV_API_KEY)
	if apiKey == "" {
		return fmt.Errorf("Auth: You must set env var %s", ENV_API_KEY)
	}

	c.Header.Add("Authorization", "Token "+apiKey)
	return nil
}
