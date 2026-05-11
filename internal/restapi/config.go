package restapi

import (
	"net/http"
	"time"
)

type Config struct {
	BaseURL string
	HTTP    *http.Client
	Header  http.Header
}

func DefaultConfig() *Config {
	return &Config{
		BaseURL: "",
		HTTP: &http.Client{
			Timeout: 30 * time.Second,
		},
		Header: make(http.Header),
	}
}

func (c *Config) SkipTLSVerify() {
	SkipTLSVerify(c.HTTP)
}
