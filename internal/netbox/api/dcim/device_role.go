package dcim

import (
	"github.com/filprpx/eg-gotomation/internal/netbox/api"
	"github.com/filprpx/eg-gotomation/internal/netbox/models"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type DeviceRoleAPI struct {
	*api.BaseAPI[models.DeviceRole]
}

func NewDeviceRoleAPI(d restapi.Doer) *DeviceRoleAPI {
	return &DeviceRoleAPI{
		BaseAPI: api.NewBaseAPI[models.DeviceRole](d, "/api/dcim/device-roles/"),
	}
}
