package ipam

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	ipamEndpointVLAN = "ipam/vlans/"
)

// VLANGet : Get a VLAN by UUID identifier.
func (c *Client) VLANGet(id uuid.UUID) (*types.VLAN, error) {
	return core.Get[types.VLAN](c.Client, ipamEndpointVLAN, id)
}

// VLANFilter : Get a list of VLANs based on query parameters.
func (c *Client) VLANFilter(q *url.Values) ([]types.VLAN, error) {
	vlans := make([]types.VLAN, 0)
	return vlans, core.Paginate[types.VLAN](c.Client, ipamEndpointVLAN, q, &vlans)
}

// VLANAll : Get all VLANs in Nautobot.
func (c *Client) VLANAll() ([]types.VLAN, error) {
	vlans := make([]types.VLAN, 0)
	return vlans, core.Paginate[types.VLAN](c.Client, ipamEndpointVLAN, nil, &vlans)
}

// VLANCreate : Generate a new VLAN record in Nautobot.
func (c *Client) VLANCreate(obj types.NewVLAN) (*types.VLAN, error) {
	return core.Create[types.VLAN, types.NewVLAN](c.Client, ipamEndpointVLAN, obj)
}

// VLANDelete : Delete a VLAN by UUID identifier.
func (c *Client) VLANDelete(id uuid.UUID) error {
	return core.Delete(c.Client, ipamEndpointVLAN, id)
}

// VLANUpdate : Update an existing VLAN record in Nautobot.
func (c *Client) VLANUpdate(id uuid.UUID, patch map[string]any) (*types.VLAN, error) {
	return core.Update[types.VLAN](c.Client, ipamEndpointVLAN, id, patch)
}
