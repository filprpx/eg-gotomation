package dcim

import (
	"github.com/filprpx/eg-gotomation/internal/netbox/api"
	"github.com/filprpx/eg-gotomation/internal/netbox/models"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type DeviceTypeAPI struct {
	*api.BaseAPI[models.DeviceType]
}

func NewDeviceTypeAPI(d restapi.Doer) *DeviceTypeAPI {
	return &DeviceTypeAPI{
		BaseAPI: api.NewBaseAPI[models.DeviceType](d, "/api/dcim/device-types/"),
	}
}
