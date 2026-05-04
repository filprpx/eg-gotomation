package dcim

import (
	"github.com/filprpx/eg-gotomation/internal/netbox/api"
	"github.com/filprpx/eg-gotomation/internal/netbox/models"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type CableAPI struct {
	*api.BaseAPI[models.Cable]
}

func NewCableAPI(d restapi.Doer) *CableAPI {
	return &CableAPI{
		BaseAPI: api.NewBaseAPI[models.Cable](d, "/api/dcim/cables/"),
	}
}
