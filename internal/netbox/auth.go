package netbox

const ENV_API_KEY = "NETBOX_API_KEY"

type Auth struct {
	ApiKey string
}

func NewAuth(apiKey string) *Auth {
	auth := Auth{
		ApiKey: apiKey,
	}

	return &auth
}
