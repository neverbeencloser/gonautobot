package extras

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
)

const (
	extrasEndpointStatus = "extras/statuses/"
)

type (
	// Status : Data type entry for a Status in Nautobot.
	Status struct {
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
		URL          string         `json:"url"`
	}

	// NewStatus : Structured input for a new Status record in Nautobot.
	NewStatus struct {
		Name         string         `json:"name"`
		ContentTypes []string       `json:"content_types"`
		Color        string         `json:"color,omitempty"`
		Description  string         `json:"description,omitempty"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
	}
)

// StatusGet : Go function to process requests for the endpoint: /api/extras/statuses/:id/
func (c *Client) StatusGet(id uuid.UUID) (*Status, error) {
	return core.Get[Status](c.Client, extrasEndpointStatus, id)
}

// StatusFilter : Go function to process requests for the endpoint: /api/extras/statuses/
func (c *Client) StatusFilter(q *url.Values) ([]Status, error) {
	statuses := make([]Status, 0)
	return statuses, core.Paginate[Status](c.Client, "extras/statuses/", q, &statuses)
}

// StatusAll : Go function to process requests for the endpoint: /api/extras/statuses/
func (c *Client) StatusAll() ([]Status, error) {
	statuses := make([]Status, 0)
	return statuses, core.Paginate[Status](c.Client, "extras/statuses/", nil, &statuses)
}

// StatusCreate : Generate a new Status record in Nautobot.
func (c *Client) StatusCreate(obj NewStatus) (*Status, error) {
	return core.Create[Status, NewStatus](c.Client, extrasEndpointStatus, obj)
}

// StatusDelete : Delete a Status by UUID identifier.
func (c *Client) StatusDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointStatus, id)
}

// StatusUpdate : Update an existing Status record in Nautobot.
func (c *Client) StatusUpdate(id uuid.UUID, patch map[string]any) (*Status, error) {
	return core.Update[Status](c.Client, extrasEndpointStatus, id, patch)
}
