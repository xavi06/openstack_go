package openstackgo

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

// Conn func
// 连接openstack集群
func Conn(endpoint string, username string, password string, tenantid string, domain string, region string) (client *gophercloud.ServiceClient, err error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: endpoint,
		Username:         username,
		Password:         password,
		TenantID:         tenantid,
		DomainID:         domain,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return
	}
	client, err = openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
	if err != nil {
		return
	}
	return
}
