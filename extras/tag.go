package extras

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/types"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
)

const (
	// extrasEndpointTag : Endpoint for interacting with Tags in Nautobot.
	extrasEndpointTag = "extras/tags/"
)

// TagGet : Fetches a tag from Nautobot by the slug name.
func (c *Client) TagGet(id uuid.UUID, opts ...query.Option) (*types.Tag, error) {
	return core.Get[types.Tag](c.Client, extrasEndpointTag, id, opts...)
}

// TagList : Get a list of Tags in Nautobot, filtered by any query options.
func (c *Client) TagList(opts ...query.Option) ([]types.Tag, error) {
	resp := make([]types.Tag, 0)
	return resp, core.Paginate[types.Tag](c.Client, extrasEndpointTag, opts, &resp)
}

// TagCreate : Generate a new Tag in Nautobot.
func (c *Client) TagCreate(tag types.NewTag) (*types.Tag, error) {
	return core.Create[types.Tag, types.NewTag](c.Client, extrasEndpointTag, tag)
}

// TagDelete : Extras method to delete a tag by UUIDv4 identifier.
func (c *Client) TagDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointTag, id)
}

// TagUpdate : Update an existing Tag in Nautobot.
func (c *Client) TagUpdate(id uuid.UUID, patch map[string]any) (*types.Tag, error) {
	return core.Update[types.Tag](c.Client, extrasEndpointTag, id, patch)
}
