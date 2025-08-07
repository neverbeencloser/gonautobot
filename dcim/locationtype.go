package dcim

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
)

const (
	dcimEndpointLocationType = "dcim/location-types/"
)

type (
	// LocationType : Represents a location type in Nautobot.
	LocationType struct {
		ID           uuid.UUID      `json:"id"`
		ContentTypes []string       `json:"content_types"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		Nestable     bool           `json:"nestable"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		Parent       *LocationType  `json:"parent"`
		TreeDepth    *int           `json:"tree_depth"`
		URL          string         `json:"url"`
	}

	// NewLocationType : Represents a new location type to be created in Nautobot.
	NewLocationType struct {
		Name         string         `json:"name"`
		ContentTypes []string       `json:"content_types"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
		Description  string         `json:"description,omitempty"`
		Nestable     bool           `json:"nestable,omitempty"`
		Parent       string         `json:"parent,omitempty"`
	}
)

// LocationTypeGet : Get a LocationType by UUID identifier.
func (c *Client) LocationTypeGet(id uuid.UUID) (*LocationType, error) {
	return core.Get[LocationType](c.Client, dcimEndpointLocationType, id)
}

// LocationTypeFilter : Get a list of LocationTypes based on query parameters.
func (c *Client) LocationTypeFilter(q *url.Values) ([]LocationType, error) {
	locationTypes := make([]LocationType, 0)
	return locationTypes, core.Paginate[LocationType](c.Client, dcimEndpointLocationType, q, &locationTypes)
}

// LocationTypeAll : Get all LocationTypes in Nautobot.
func (c *Client) LocationTypeAll() ([]LocationType, error) {
	locationTypes := make([]LocationType, 0)
	return locationTypes, core.Paginate[LocationType](c.Client, dcimEndpointLocationType, nil, &locationTypes)
}

// LocationTypeCreate : Generate a new LocationType record in Nautobot.
func (c *Client) LocationTypeCreate(obj NewLocationType) (*LocationType, error) {
	return core.Create[LocationType, NewLocationType](c.Client, dcimEndpointLocationType, obj)
}

// LocationTypeUpdate : Update an existing LocationType record in Nautobot.
func (c *Client) LocationTypeUpdate(id uuid.UUID, patch map[string]any) (*LocationType, error) {
	return core.Update[LocationType](c.Client, dcimEndpointLocationType, id, patch)
}

// LocationTypeDelete : Delete a LocationType by UUID identifier.
func (c *Client) LocationTypeDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointLocationType, id)
}
