package ipam

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	ipamEndpointVRF = "ipam/vrfs/"
)

// VRFGet : Get a VRF by UUID identifier.
func (c *Client) VRFGet(id uuid.UUID, opts ...query.Option) (*types.VRF, error) {
	return core.Get[types.VRF](c.Client, ipamEndpointVRF, id, opts...)
}

// VRFList : Get a list of VRFs in Nautobot, filtered by any query options.
func (c *Client) VRFList(opts ...query.Option) ([]types.VRF, error) {
	vrfs := make([]types.VRF, 0)
	return vrfs, core.Paginate[types.VRF](c.Client, ipamEndpointVRF, opts, &vrfs)
}

// VRFCreate : Generate a new VRF record in Nautobot.
func (c *Client) VRFCreate(obj types.NewVRF) (*types.VRF, error) {
	return core.Create[types.VRF, types.NewVRF](c.Client, ipamEndpointVRF, obj)
}

// VRFDelete : Delete a VRF by UUID identifier.
func (c *Client) VRFDelete(id uuid.UUID) error {
	return core.Delete(c.Client, ipamEndpointVRF, id)
}

// VRFUpdate : Update an existing VRF record in Nautobot.
func (c *Client) VRFUpdate(id uuid.UUID, patch map[string]any) (*types.VRF, error) {
	return core.Update[types.VRF](c.Client, ipamEndpointVRF, id, patch)
}
