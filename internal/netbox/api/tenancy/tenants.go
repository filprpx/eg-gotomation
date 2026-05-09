package tenancy

import (
	"github.com/filprpx/eg-gotomation/internal/netbox/api"
	"github.com/filprpx/eg-gotomation/internal/netbox/models"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type TenantAPI struct {
	*api.BaseAPI[models.Tenant]
}

func NewTenantAPI(d restapi.Doer) *TenantAPI {
	return &TenantAPI{
		BaseAPI: api.NewBaseAPI[models.Tenant](d, "/api/dcim/tenants/"),
	}
}
