package netbox

import ()

type DeviceRolesAPI struct {
	client *NetboxClient
}

type DeviceRole struct {
	ApiBaseFields
}

type DeviceRoleWrite struct {
	DeviceRole
}
