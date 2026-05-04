package netbox

import (
	"fmt"
	"os"

	"github.com/filprpx/eg-gotomation/internal/restapi"
)

const ENV_API_BASE_URL = "NETBOX_API_BASE_URL"

type Config struct {
	restapi.Config
	Auth
}

func DefaultConfig() *Config {
	return &Config{
		Config: *restapi.NewConfig(),
		Auth:   *NewAuth(""),
	}
}

func ConfigFromEnv() (*Config, error) {
	baseUrl := os.Getenv(ENV_API_BASE_URL)
	if baseUrl == "" {
		return nil, fmt.Errorf("ConfigFromEnv: %s not configured", ENV_API_BASE_URL)
	}

	apiKey := os.Getenv(ENV_API_KEY)
	if apiKey == "" {
		return nil, fmt.Errorf("ConfigFromEnv: %s not configured", ENV_API_KEY)
	}

	cfg := DefaultConfig()
	cfg.BaseUrl = baseUrl
	cfg.ApiKey = apiKey

	return cfg, nil
}

func ValidateConfig(cfg *Config) error {
	if cfg.BaseUrl == "" {
		return fmt.Errorf("BaseUrl must be specified")
	}

	if cfg.ApiKey == "" {
		return fmt.Errorf("ApiKey must be specified")
	}

	return nil
}
