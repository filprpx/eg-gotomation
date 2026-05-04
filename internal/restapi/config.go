package restapi

import (
	"net/http"
	"time"
)

type Config struct {
	BaseUrl string
	HTTP    *http.Client
	Header  http.Header
}

func NewConfig() *Config {
	return &Config{
		BaseUrl: "",
		HTTP: &http.Client{
			Timeout: 30 * time.Second,
		},
		Header: make(http.Header),
	}
}

func (c *Config) SkipTLSVerify() {
	SkipTLSVerify(c.HTTP)
}
