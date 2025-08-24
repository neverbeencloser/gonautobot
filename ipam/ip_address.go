package ipam

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	ipamEndpointIPAddress = "ipam/ip-addresses/"
)

// IPAddressGet : Get an IP Address by UUID identifier.
func (c *Client) IPAddressGet(id uuid.UUID) (*types.IPAddress, error) {
	return core.Get[types.IPAddress](c.Client, ipamEndpointIPAddress, id)
}

// IPAddressFilter : Get a list of IP Addresses based on query parameters.
func (c *Client) IPAddressFilter(q *url.Values) ([]types.IPAddress, error) {
	ipAddresses := make([]types.IPAddress, 0)
	return ipAddresses, core.Paginate[types.IPAddress](c.Client, ipamEndpointIPAddress, q, &ipAddresses)
}

// IPAddressAll : Get all IP Addresses in Nautobot.
func (c *Client) IPAddressAll() ([]types.IPAddress, error) {
	ipAddresses := make([]types.IPAddress, 0)
	return ipAddresses, core.Paginate[types.IPAddress](c.Client, ipamEndpointIPAddress, nil, &ipAddresses)
}

// IPAddressCreate : Generate a new IP Address record in Nautobot.
func (c *Client) IPAddressCreate(obj types.NewIPAddress) (*types.IPAddress, error) {
	return core.Create[types.IPAddress, types.NewIPAddress](c.Client, ipamEndpointIPAddress, obj)
}

// IPAddressDelete : Delete an IP Address by UUID identifier.
func (c *Client) IPAddressDelete(id uuid.UUID) error {
	return core.Delete(c.Client, ipamEndpointIPAddress, id)
}

// IPAddressUpdate : Update an existing IP Address record in Nautobot.
func (c *Client) IPAddressUpdate(id uuid.UUID, patch map[string]any) (*types.IPAddress, error) {
	return core.Update[types.IPAddress](c.Client, ipamEndpointIPAddress, id, patch)
}
