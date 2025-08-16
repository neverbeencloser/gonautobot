package dcim

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointLocationType = "dcim/location-types/"
)

// LocationTypeGet : Get a LocationType by UUID identifier.
func (c *Client) LocationTypeGet(id uuid.UUID) (*types.LocationType, error) {
	return core.Get[types.LocationType](c.Client, dcimEndpointLocationType, id)
}

// LocationTypeFilter : Get a list of LocationTypes based on query parameters.
func (c *Client) LocationTypeFilter(q *url.Values) ([]types.LocationType, error) {
	locationTypes := make([]types.LocationType, 0)
	return locationTypes, core.Paginate[types.LocationType](c.Client, dcimEndpointLocationType, q, &locationTypes)
}

// LocationTypeAll : Get all LocationTypes in Nautobot.
func (c *Client) LocationTypeAll() ([]types.LocationType, error) {
	locationTypes := make([]types.LocationType, 0)
	return locationTypes, core.Paginate[types.LocationType](c.Client, dcimEndpointLocationType, nil, &locationTypes)
}

// LocationTypeCreate : Generate a new LocationType record in Nautobot.
func (c *Client) LocationTypeCreate(obj types.NewLocationType) (*types.LocationType, error) {
	return core.Create[types.LocationType, types.NewLocationType](c.Client, dcimEndpointLocationType, obj)
}

// LocationTypeUpdate : Update an existing LocationType record in Nautobot.
func (c *Client) LocationTypeUpdate(id uuid.UUID, patch map[string]any) (*types.LocationType, error) {
	return core.Update[types.LocationType](c.Client, dcimEndpointLocationType, id, patch)
}

// LocationTypeDelete : Delete a LocationType by UUID identifier.
func (c *Client) LocationTypeDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointLocationType, id)
}
