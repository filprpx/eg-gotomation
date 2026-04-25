package netbox

import ()

type SitesAPI struct {
	client *NetboxClient
}

type Site struct {
	ApiBaseFields
}

type SiteWrite struct {
	Site
}
