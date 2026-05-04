package dcim

import (
	"github.com/filprpx/eg-gotomation/internal/netbox/api"
	"github.com/filprpx/eg-gotomation/internal/netbox/models"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type ManufacturerAPI struct {
	*api.BaseAPI[models.Manufacturer]
}

func NewManufacturerAPI(d restapi.Doer) *ManufacturerAPI {
	return &ManufacturerAPI{
		BaseAPI: api.NewBaseAPI[models.Manufacturer](d, "/api/dcim/manufacturers/"),
	}
}
