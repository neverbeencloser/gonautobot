package extras

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	extrasEndpointStatus = "extras/statuses/"
)

// StatusGet : Go function to process requests for the endpoint: /api/extras/statuses/:id/
func (c *Client) StatusGet(id uuid.UUID) (*types.Status, error) {
	return core.Get[types.Status](c.Client, extrasEndpointStatus, id)
}

// StatusFilter : Go function to process requests for the endpoint: /api/extras/statuses/
func (c *Client) StatusFilter(q *url.Values) ([]types.Status, error) {
	statuses := make([]types.Status, 0)
	return statuses, core.Paginate[types.Status](c.Client, "extras/statuses/", q, &statuses)
}

// StatusAll : Go function to process requests for the endpoint: /api/extras/statuses/
func (c *Client) StatusAll() ([]types.Status, error) {
	statuses := make([]types.Status, 0)
	return statuses, core.Paginate[types.Status](c.Client, "extras/statuses/", nil, &statuses)
}

// StatusCreate : Generate a new Status record in Nautobot.
func (c *Client) StatusCreate(obj types.NewStatus) (*types.Status, error) {
	return core.Create[types.Status, types.NewStatus](c.Client, extrasEndpointStatus, obj)
}

// StatusDelete : Delete a Status by UUID identifier.
func (c *Client) StatusDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointStatus, id)
}

// StatusUpdate : Update an existing Status record in Nautobot.
func (c *Client) StatusUpdate(id uuid.UUID, patch map[string]any) (*types.Status, error) {
	return core.Update[types.Status](c.Client, extrasEndpointStatus, id, patch)
}
