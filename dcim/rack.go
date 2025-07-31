package dcim

import (
	"net/url"
	"time"

	"github.com/josh-silvas/gonautobot/core"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared"
	"github.com/josh-silvas/gonautobot/tenancy"

	"github.com/google/uuid"
)

const (
	dcimEndpointRack = "dcim/racks/"
)

type (
	// Rack : Data type entry for a Rack in Nautobot.
	Rack struct {
		ID             uuid.UUID            `json:"id"`
		AssetTag       *string              `json:"asset_tag"`
		Comments       string               `json:"comments"`
		Created        time.Time            `json:"created"`
		CustomFields   map[string]any       `json:"custom_fields"`
		DescUnits      bool                 `json:"desc_units"`
		DeviceCount    int                  `json:"device_count"`
		Display        string               `json:"display"`
		FacilityID     *string              `json:"facility_id"`
		LastUpdated    time.Time            `json:"last_updated"`
		Location       Location             `json:"location"`
		Name           string               `json:"name"`
		NaturalSlug    string               `json:"natural_slug"`
		NotesURL       string               `json:"notes_url"`
		ObjectType     string               `json:"object_type"`
		OuterDepth     *int                 `json:"outer_depth"`
		OuterUnit      *shared.LabelValue   `json:"outer_unit"`
		OuterWidth     *int                 `json:"outer_width"`
		PowerFeedCount int                  `json:"power_feed_count"`
		RackGroup      *RackGroup           `json:"rack_group"`
		Role           *extras.Role         `json:"role"`
		Serial         string               `json:"serial"`
		Status         extras.Status        `json:"status"`
		Tags           []extras.Tag         `json:"tags"`
		Tenant         *tenancy.Tenant      `json:"tenant"`
		Type           *shared.LabelValue   `json:"type"`
		UHeight        int                  `json:"u_height"`
		URL            string               `json:"url"`
		Width          shared.LabelValueInt `json:"width"`
	}

	// NewRack represents the structure for creating a new Rack in Nautobot.
	NewRack struct {
		Location      string         `json:"location"`
		Name          string         `json:"name"`
		Status        string         `json:"status"`
		Width         int            `json:"width"`
		UHeight       int            `json:"u_height"`
		AssetTag      string         `json:"asset_tag,omitempty"`
		Comments      string         `json:"comments,omitempty"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		DescUnits     bool           `json:"desc_units,omitempty"`
		FacilityID    string         `json:"facility_id,omitempty"`
		OuterDepth    int            `json:"outer_depth,omitempty"`
		OuterUnit     string         `json:"outer_unit,omitempty"`
		OuterWidth    int            `json:"outer_width,omitempty"`
		RackGroup     string         `json:"rack_group,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
		Role          string         `json:"role,omitempty"`
		Serial        string         `json:"serial,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
		Tenant        string         `json:"tenant,omitempty"`
		Type          string         `json:"type,omitempty"`
	}
)

// RackGet : Get a Rack by UUID identifier.
func (c *Client) RackGet(id uuid.UUID) (*Rack, error) {
	return core.Get[Rack](c.Client, dcimEndpointRack, id)
}

// RackFilter : Get a list of Racks based on query parameters.
func (c *Client) RackFilter(q *url.Values) ([]Rack, error) {
	racks := make([]Rack, 0)
	return racks, core.Paginate[Rack](c.Client, dcimEndpointRack, q, &racks)
}

// RackAll : Get all Racks in Nautobot.
func (c *Client) RackAll() ([]Rack, error) {
	racks := make([]Rack, 0)
	return racks, core.Paginate[Rack](c.Client, dcimEndpointRack, nil, &racks)
}

// RackCreate : Generate a new Rack record in Nautobot.
func (c *Client) RackCreate(obj NewRack) (*Rack, error) {
	return core.Create[Rack, NewRack](c.Client, dcimEndpointRack, obj)
}

// RackDelete : Delete a Rack by UUID identifier.
func (c *Client) RackDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointRack, id)
}

// RackUpdate : Update an existing Rack record in Nautobot.
func (c *Client) RackUpdate(id uuid.UUID, patch map[string]any) (*Rack, error) {
	return core.Update[Rack](c.Client, dcimEndpointRack, id, patch)
}
