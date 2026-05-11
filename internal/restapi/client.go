// Package restapi is used as basis of common definitions and actions any concrete client that interacts with rest apis may need
package restapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Config
}

func NewClient(baseURL string) *Client {
	cfg := DefaultConfig()
	cfg.BaseURL = baseURL

	return NewClientWithConfig(cfg)
}

func NewClientWithConfig(cfg *Config) *Client {
	return &Client{
		Config: *cfg,
	}
}

func (c *Client) Do(ctx context.Context, req Request) (*http.Response, error) {
	url := c.BaseURL + req.Path
	if req.Query != nil {
		url = url + "?" + req.Query.Encode()
	}

	var body io.Reader
	if req.Body != nil {
		b, err := json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("marshal: %w", err)
		}
		body = bytes.NewReader(b)
	}

	httpReq, err := http.NewRequestWithContext(ctx, req.Method, url, body)
	if err != nil {
		return nil, err
	}

	for key, values := range c.Header {
		for _, value := range values {
			httpReq.Header.Add(key, value)
		}
	}

	if body != nil {
		httpReq.Header.Add("Content-Type", "application/json")
	}

	return c.HTTP.Do(httpReq)
}
