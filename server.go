package openstackgo

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

// GetServers func
func GetServers(client *gophercloud.ServiceClient) (allServers []servers.Server, err error) {
	listOpts := servers.ListOpts{
		AllTenants: false,
	}

	allPages, err := servers.List(client, listOpts).AllPages()
	if err != nil {
		return
	}

	allServers, err = servers.ExtractServers(allPages)
	if err != nil {
		return
	}

	return
}
