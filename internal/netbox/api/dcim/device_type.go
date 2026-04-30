package dcim

import "github.com/filprpx/eg-gotomation/internal/restapi"

type DeviceTypeAPI struct {
	doer *restapi.Doer
}

func NewDeviceTypeAPI(d *restapi.Doer) *DeviceTypeAPI {
	return &DeviceTypeAPI{
		doer: d,
	}
}
