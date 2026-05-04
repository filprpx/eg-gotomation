package dcim

import (
	"github.com/filprpx/eg-gotomation/internal/netbox/api"
	"github.com/filprpx/eg-gotomation/internal/netbox/models"
	"github.com/filprpx/eg-gotomation/internal/restapi"
)

type SiteAPI struct {
	*api.BaseAPI[models.Site]
}

func NewSiteAPI(d restapi.Doer) *SiteAPI {
	return &SiteAPI{
		BaseAPI: api.NewBaseAPI[models.Site](d, "/api/dcim/sites/"),
	}
}
