package dcim

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointPlatform = "dcim/platforms/"
)

// PlatformGet : Get a Platform by UUID identifier.
func (c *Client) PlatformGet(id uuid.UUID, opts ...query.Option) (*types.Platform, error) {
	return core.Get[types.Platform](c.Client, dcimEndpointPlatform, id, opts...)
}

// PlatformList : Get a list of Platforms in Nautobot, filtered by any query options.
func (c *Client) PlatformList(opts ...query.Option) ([]types.Platform, error) {
	platforms := make([]types.Platform, 0)
	return platforms, core.Paginate[types.Platform](c.Client, dcimEndpointPlatform, opts, &platforms)
}

// PlatformCreate : Generate a new Platform record in Nautobot.
func (c *Client) PlatformCreate(obj types.NewPlatform) (*types.Platform, error) {
	return core.Create[types.Platform, types.NewPlatform](c.Client, dcimEndpointPlatform, obj)
}

// PlatformDelete : Delete a Platform by UUID identifier.
func (c *Client) PlatformDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointPlatform, id)
}

// PlatformUpdate : Update an existing Platform record in Nautobot.
func (c *Client) PlatformUpdate(id uuid.UUID, patch map[string]any) (*types.Platform, error) {
	return core.Update[types.Platform](c.Client, dcimEndpointPlatform, id, patch)
}
