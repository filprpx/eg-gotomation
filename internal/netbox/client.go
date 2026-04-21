package netbox

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const ENV_URL = "NETBOX_URL"

type NetboxClient struct {
	url    string
	http   *http.Client
	header http.Header
	auth   *NetboxAuth

	// api definitions
	DCIM *DcimAPI
	IPAM *IpamAPI
}

func NewClient() (*NetboxClient, error) {
	url := os.Getenv(ENV_URL)
	if url == "" {
		return nil, fmt.Errorf("%s not configured", ENV_URL)
	}

	var auth, err = NewAuth()
	if err != nil {
		return nil, err
	}

	client := &NetboxClient{
		auth: auth,
		url:  url,
		http: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	client.Auth()

	client.DCIM, _ = NewDcimAPI(client)
	client.IPAM, _ = NewIpamAPI(client)

	return client, nil
}

func (c *NetboxClient) Auth() {
	c.header = make(http.Header)
	c.header.Add("Authorization", "Token "+c.auth.APIKey)
	c.header.Add("Accept", "application/json")
}

func (c *NetboxClient) setHeader(req *http.Request) {
	for key, values := range c.header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
}

func (c *NetboxClient) get(ctx context.Context, path string) (*http.Response, error) {
	url := c.url + path
	req, _ := http.NewRequestWithContext(ctx, "get", url, nil)
	c.setHeader(req)
	return c.http.Do(req)
}

func (c *NetboxClient) post(ctx context.Context, path string) (*http.Response, error) {
	url := c.url + path
	req, _ := http.NewRequestWithContext(ctx, "post", url, nil)
	c.setHeader(req)
	return c.http.Do(req)
}

func (c *NetboxClient) patch(ctx context.Context, path string) (*http.Response, error) {
	url := c.url + path
	req, _ := http.NewRequestWithContext(ctx, "patch", url, nil)
	c.setHeader(req)
	return c.http.Do(req)
}

func (c *NetboxClient) delete(ctx context.Context, path string) (*http.Response, error) {
	url := c.url + path
	req, _ := http.NewRequestWithContext(ctx, "delete", url, nil)
	c.setHeader(req)
	return c.http.Do(req)
}

func (c *NetboxClient) apiError(r *http.Response) error {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return fmt.Errorf("netbox api error: %d (failed to read body %v)", r.StatusCode, err)
	}

	return fmt.Errorf("netbox api error : %d: %s", r.StatusCode, body)
}
