package ipam

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	ipamEndpointVLAN = "ipam/vlans/"
)

// VLANGet : Get a VLAN by UUID identifier.
func (c *Client) VLANGet(id uuid.UUID, opts ...query.Option) (*types.VLAN, error) {
	return core.Get[types.VLAN](c.Client, ipamEndpointVLAN, id, opts...)
}

// VLANList : Get a list of VLANs in Nautobot, filtered by any query options.
func (c *Client) VLANList(opts ...query.Option) ([]types.VLAN, error) {
	vlans := make([]types.VLAN, 0)
	return vlans, core.Paginate[types.VLAN](c.Client, ipamEndpointVLAN, opts, &vlans)
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
