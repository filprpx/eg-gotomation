package netbox

import (
	"fmt"
	"os"

	"github.com/filprpx/eg-gotomation/internal/restapi"
)

const EnvAPIBaseURL = "NETBOX_API_BASE_URL"

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
	baseURL := os.Getenv(EnvAPIBaseURL)
	if baseURL == "" {
		return nil, fmt.Errorf("ConfigFromEnv: %s not configured", EnvAPIBaseURL)
	}

	apiKey := os.Getenv(EnvAPIKey)
	if apiKey == "" {
		return nil, fmt.Errorf("ConfigFromEnv: %s not configured", EnvAPIKey)
	}

	cfg := DefaultConfig()
	cfg.BaseURL = baseURL
	cfg.APIKey = apiKey

	return cfg, nil
}

func ValidateConfig(cfg *Config) error {
	if cfg.BaseURL == "" {
		return fmt.Errorf("BaseUrl must be specified")
	}

	if cfg.APIKey == "" {
		return fmt.Errorf("ApiKey must be specified")
	}

	return nil
}
