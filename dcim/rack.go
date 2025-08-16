package dcim

import (
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"

	"github.com/google/uuid"
)

const (
	dcimEndpointRack = "dcim/racks/"
)

// RackGet : Get a Rack by UUID identifier.
func (c *Client) RackGet(id uuid.UUID) (*types.Rack, error) {
	return core.Get[types.Rack](c.Client, dcimEndpointRack, id)
}

// RackFilter : Get a list of Racks based on query parameters.
func (c *Client) RackFilter(q *url.Values) ([]types.Rack, error) {
	racks := make([]types.Rack, 0)
	return racks, core.Paginate[types.Rack](c.Client, dcimEndpointRack, q, &racks)
}

// RackAll : Get all Racks in Nautobot.
func (c *Client) RackAll() ([]types.Rack, error) {
	racks := make([]types.Rack, 0)
	return racks, core.Paginate[types.Rack](c.Client, dcimEndpointRack, nil, &racks)
}

// RackCreate : Generate a new Rack record in Nautobot.
func (c *Client) RackCreate(obj types.NewRack) (*types.Rack, error) {
	return core.Create[types.Rack, types.NewRack](c.Client, dcimEndpointRack, obj)
}

// RackDelete : Delete a Rack by UUID identifier.
func (c *Client) RackDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointRack, id)
}

// RackUpdate : Update an existing Rack record in Nautobot.
func (c *Client) RackUpdate(id uuid.UUID, patch map[string]any) (*types.Rack, error) {
	return core.Update[types.Rack](c.Client, dcimEndpointRack, id, patch)
}
