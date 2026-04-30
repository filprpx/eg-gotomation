package restapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

type BaseClient struct {
	BaseUrl string
	HTTP    *http.Client
	Header  http.Header
}

func NewBaseClient(baseUrl string) *BaseClient {
	return &BaseClient{
		BaseUrl: baseUrl,
		HTTP: &http.Client{
			Timeout: 30 * time.Second,
		},
		Header: make(http.Header),
	}
}

func (c *BaseClient) Get(ctx context.Context, path string) (*http.Response, error) {
	return c.do(ctx, "GET", path, nil)
}

func (c *BaseClient) Post(ctx context.Context, path string, body io.Reader) (*http.Response, error) {
	return c.do(ctx, "POST", path, body)
}

func (c *BaseClient) Patch(ctx context.Context, path string, body io.Reader) (*http.Response, error) {
	return c.do(ctx, "PATCH", path, body)
}

func (c *BaseClient) Delete(ctx context.Context, path string) (*http.Response, error) {
	return c.do(ctx, "DELETE", path, nil)
}

func (c *BaseClient) do(ctx context.Context, method string, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.BaseUrl+path, body)
	if err != nil {
		return nil, err
	}

	for key, values := range c.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return c.HTTP.Do(req)
}
