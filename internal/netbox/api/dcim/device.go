package dcim

import (
	"github.com/filprpx/eg-gotomation/internal/netbox/api"
	"github.com/filprpx/eg-gotomation/internal/netbox/models"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type DeviceAPI struct {
	*api.BaseAPI[models.Device]
}

func NewDeviceAPI(d restapi.Doer) *DeviceAPI {
	return &DeviceAPI{
		BaseAPI: api.NewBaseAPI[models.Device](d, "/api/dcim/devices/"),
	}
}

// func (d *DeviceAPI) List(ctx context.Context) (*[]models.Device, error) {
// 	return &[]models.Device{}, nil
// }
//
// func (d *DeviceAPI) Get(ctx context.Context, id int) (*models.Device, error) {
// 	return &models.Device{}, nil
// }
//
// func (d *DeviceAPI) Create(ctx context.Context, device *models.Device) (*models.Device, error) {
// 	return &models.Device{}, nil
// }
//
// func (d *DeviceAPI) Update(ctx context.Context, device *models.Device) (bool, error) {
// 	return true, nil
// }
//
// func (d *DeviceAPI) Delete(ctx context.Context, id int) (bool, error) {
// 	return true, nil
// }
