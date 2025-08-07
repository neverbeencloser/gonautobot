package extras

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/neverbeencloser/gonautobot/core"
)

const (
	// extrasEndpointTag : Endpoint for interacting with Tags in Nautobot.
	extrasEndpointTag = "extras/tags/"
)

type (
	// Tag : Data type entry for a Tag in Nautobot.
	Tag struct {
		ID           uuid.UUID      `json:"id"`
		Color        string         `json:"color"`
		ContentTypes []string       `json:"content_types"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		TaggedItems  int            `json:"tagged_items"`
		URL          string         `json:"url"`
	}

	// NewTag : Data type for creating a new Tag in Nautobot.
	NewTag struct {
		Name          string         `json:"name"`
		ContentTypes  []string       `json:"content_types"`
		Color         string         `json:"color,omitempty"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
	}
)

// TagGet : Fetches a tag from Nautobot by the slug name.
func (c *Client) TagGet(id uuid.UUID) (*Tag, error) {
	return core.Get[Tag](c.Client, extrasEndpointTag, id)
}

// TagFilter : Returns an array of tags from Nautobot filtered by the q (query values)
func (c *Client) TagFilter(q *url.Values) ([]Tag, error) {
	resp := make([]Tag, 0)
	return resp, core.Paginate[Tag](c.Client, extrasEndpointTag, q, &resp)
}

// TagAll : Returns an array of all tags from Nautobot
func (c *Client) TagAll() ([]Tag, error) {
	resp := make([]Tag, 0)
	return resp, core.Paginate[Tag](c.Client, extrasEndpointTag, nil, &resp)
}

// TagCreate : Generate a new Tag in Nautobot.
func (c *Client) TagCreate(tag NewTag) (*Tag, error) {
	return core.Create[Tag, NewTag](c.Client, extrasEndpointTag, tag)
}

// TagDelete : Extras method to delete a tag by UUIDv4 identifier.
func (c *Client) TagDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointTag, id)
}

// TagUpdate : Update an existing Tag in Nautobot.
func (c *Client) TagUpdate(id uuid.UUID, patch map[string]any) (*Tag, error) {
	return core.Update[Tag](c.Client, extrasEndpointTag, id, patch)
}
