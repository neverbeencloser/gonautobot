package dcim

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointRackGroup = "dcim/rack-groups/"
)

// RackGroupGet : Get a RackGroup by UUID identifier.
func (c *Client) RackGroupGet(id uuid.UUID) (*types.RackGroup, error) {
	return core.Get[types.RackGroup](c.Client, dcimEndpointRackGroup, id)
}

// RackGroupFilter : Get a list of RackGroups based on query parameters.
func (c *Client) RackGroupFilter(q *url.Values) ([]types.RackGroup, error) {
	rackGroups := make([]types.RackGroup, 0)
	return rackGroups, core.Paginate[types.RackGroup](c.Client, dcimEndpointRackGroup, q, &rackGroups)
}

// RackGroupAll : Get all RackGroups in Nautobot.
func (c *Client) RackGroupAll() ([]types.RackGroup, error) {
	rackGroups := make([]types.RackGroup, 0)
	return rackGroups, core.Paginate[types.RackGroup](c.Client, dcimEndpointRackGroup, nil, &rackGroups)
}

// RackGroupCreate : Generate a new RackGroup record in Nautobot.
func (c *Client) RackGroupCreate(obj types.NewRackGroup) (*types.RackGroup, error) {
	return core.Create[types.RackGroup, types.NewRackGroup](c.Client, dcimEndpointRackGroup, obj)
}

// RackGroupDelete : Delete a RackGroup by UUID identifier.
func (c *Client) RackGroupDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointRackGroup, id)
}

// RackGroupUpdate : Update an existing RackGroup record in Nautobot.
func (c *Client) RackGroupUpdate(id uuid.UUID, patch map[string]any) (*types.RackGroup, error) {
	return core.Update[types.RackGroup](c.Client, dcimEndpointRackGroup, id, patch)
}
