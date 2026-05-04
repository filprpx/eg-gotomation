package netbox

const EnvAPIKey = "NETBOX_API_KEY"

type Auth struct {
	APIKey string
}

func NewAuth(apiKey string) *Auth {
	auth := Auth{
		APIKey: apiKey,
	}

	return &auth
}
