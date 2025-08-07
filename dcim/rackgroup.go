package dcim

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
)

const (
	dcimEndpointRackGroup = "dcim/rack-groups/"
)

type (
	// RackGroup : Data type entry for a RackGroup in Nautobot.
	RackGroup struct {
		ID           uuid.UUID      `json:"id"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Location     Location       `json:"location"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		Parent       *RackGroup     `json:"parent"`
		RackCount    int            `json:"rack_count"`
		TreeDepth    *int           `json:"tree_depth"`
		URL          string         `json:"url"`
	}

	// NewRackGroup represents the structure for creating a new RackGroup in Nautobot.
	NewRackGroup struct {
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		Location      string         `json:"location"`
		Name          string         `json:"name"`
		Parent        string         `json:"parent,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
	}
)

// RackGroupGet : Get a RackGroup by UUID identifier.
func (c *Client) RackGroupGet(id uuid.UUID) (*RackGroup, error) {
	return core.Get[RackGroup](c.Client, dcimEndpointRackGroup, id)
}

// RackGroupFilter : Get a list of RackGroups based on query parameters.
func (c *Client) RackGroupFilter(q *url.Values) ([]RackGroup, error) {
	rackGroups := make([]RackGroup, 0)
	return rackGroups, core.Paginate[RackGroup](c.Client, dcimEndpointRackGroup, q, &rackGroups)
}

// RackGroupAll : Get all RackGroups in Nautobot.
func (c *Client) RackGroupAll() ([]RackGroup, error) {
	rackGroups := make([]RackGroup, 0)
	return rackGroups, core.Paginate[RackGroup](c.Client, dcimEndpointRackGroup, nil, &rackGroups)
}

// RackGroupCreate : Generate a new RackGroup record in Nautobot.
func (c *Client) RackGroupCreate(obj NewRackGroup) (*RackGroup, error) {
	return core.Create[RackGroup, NewRackGroup](c.Client, dcimEndpointRackGroup, obj)
}

// RackGroupDelete : Delete a RackGroup by UUID identifier.
func (c *Client) RackGroupDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointRackGroup, id)
}

// RackGroupUpdate : Update an existing RackGroup record in Nautobot.
func (c *Client) RackGroupUpdate(id uuid.UUID, patch map[string]any) (*RackGroup, error) {
	return core.Update[RackGroup](c.Client, dcimEndpointRackGroup, id, patch)
}
