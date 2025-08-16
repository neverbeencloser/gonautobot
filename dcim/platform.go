package dcim

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointPlatform = "dcim/platforms/"
)

// PlatformGet : Get a Platform by UUID identifier.
func (c *Client) PlatformGet(id uuid.UUID) (*types.Platform, error) {
	return core.Get[types.Platform](c.Client, dcimEndpointPlatform, id)
}

// PlatformFilter : Get a list of Platforms based on query parameters.
func (c *Client) PlatformFilter(q *url.Values) ([]types.Platform, error) {
	platforms := make([]types.Platform, 0)
	return platforms, core.Paginate[types.Platform](c.Client, dcimEndpointPlatform, q, &platforms)
}

// PlatformAll : Get all Platforms in Nautobot.
func (c *Client) PlatformAll() ([]types.Platform, error) {
	platforms := make([]types.Platform, 0)
	return platforms, core.Paginate[types.Platform](c.Client, dcimEndpointPlatform, nil, &platforms)
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
