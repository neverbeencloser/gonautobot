package dcim

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/josh-silvas/gonautobot/core"
)

const (
	dcimEndpointPlatform = "dcim/platforms/"
)

type (
	// Platform represents a device platform in Nautobot.
	Platform struct {
		ID                    uuid.UUID      `json:"id"`
		Created               time.Time      `json:"created"`
		CustomFields          map[string]any `json:"custom_fields"`
		Description           string         `json:"description"`
		DeviceCount           int            `json:"device_count"`
		Display               string         `json:"display"`
		LastUpdated           time.Time      `json:"last_updated"`
		Manufacturer          *Manufacturer  `json:"manufacturer"`
		Name                  string         `json:"name"`
		NapalmArgs            any            `json:"napalm_args"`
		NapalmDriver          string         `json:"napalm_driver"`
		NaturalSlug           string         `json:"natural_slug"`
		NetworkDriver         string         `json:"network_driver"`
		NetworkDriverMappings map[string]any `json:"network_driver_mappings"`
		NotesURL              string         `json:"notes_url"`
		ObjectType            string         `json:"object_type"`
		URL                   string         `json:"url"`
		VMCount               int            `json:"virtual_machine_count"`
	}

	// NewPlatform represents the data required to create a new platform.
	NewPlatform struct {
		Name          string         `json:"name"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		Manufacturer  string         `json:"manufacturer,omitempty"`
		NapalmArgs    any            `json:"napalm_args,omitempty"`
		NapalmDriver  string         `json:"napalm_driver,omitempty"`
		NetworkDriver string         `json:"network_driver,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
	}
)

// PlatformGet : Get a Platform by UUID identifier.
func (c *Client) PlatformGet(id uuid.UUID) (*Platform, error) {
	return core.Get[Platform](c.Client, dcimEndpointPlatform, id)
}

// PlatformFilter : Get a list of Platforms based on query parameters.
func (c *Client) PlatformFilter(q *url.Values) ([]Platform, error) {
	platforms := make([]Platform, 0)
	return platforms, core.Paginate[Platform](c.Client, dcimEndpointPlatform, q, &platforms)
}

// PlatformAll : Get all Platforms in Nautobot.
func (c *Client) PlatformAll() ([]Platform, error) {
	platforms := make([]Platform, 0)
	return platforms, core.Paginate[Platform](c.Client, dcimEndpointPlatform, nil, &platforms)
}

// PlatformCreate : Generate a new Platform record in Nautobot.
func (c *Client) PlatformCreate(obj NewPlatform) (*Platform, error) {
	return core.Create[Platform, NewPlatform](c.Client, dcimEndpointPlatform, obj)
}

// PlatformDelete : Delete a Platform by UUID identifier.
func (c *Client) PlatformDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointPlatform, id)
}

// PlatformUpdate : Update an existing Platform record in Nautobot.
func (c *Client) PlatformUpdate(id uuid.UUID, patch map[string]any) (*Platform, error) {
	return core.Update[Platform](c.Client, dcimEndpointPlatform, id, patch)
}
