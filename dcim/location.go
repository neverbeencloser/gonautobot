package dcim

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointLocation = "dcim/locations/"
)

// LocationGet : Get a Location by UUID identifier.
func (c *Client) LocationGet(id uuid.UUID, opts ...query.Option) (*types.Location, error) {
	return core.Get[types.Location](c.Client, dcimEndpointLocation, id, opts...)
}

// LocationList : Get a list of Locations in Nautobot, filtered by any query options.
func (c *Client) LocationList(opts ...query.Option) ([]types.Location, error) {
	locationTypes := make([]types.Location, 0)
	return locationTypes, core.Paginate[types.Location](c.Client, dcimEndpointLocation, opts, &locationTypes)
}

// LocationCreate : Generate a new Location record in Nautobot.
func (c *Client) LocationCreate(obj types.NewLocation) (*types.Location, error) {
	return core.Create[types.Location, types.NewLocation](c.Client, dcimEndpointLocation, obj)
}

// LocationUpdate : Update an existing Location record in Nautobot.
func (c *Client) LocationUpdate(id uuid.UUID, patch map[string]any) (*types.Location, error) {
	return core.Update[types.Location](c.Client, dcimEndpointLocation, id, patch)
}

// LocationDelete : Delete a Location by UUID identifier.
func (c *Client) LocationDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointLocation, id)
}
