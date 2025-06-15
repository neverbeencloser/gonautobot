package extras

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/josh-silvas/gonautobot/core"
)

const (
	extrasEndpointRole = "extras/roles/"
)

type (
	// Role : Data type entry for a Role in Nautobot.
	Role struct {
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
		Weight       *int           `json:"weight"`
	}

	// NewRole : Structured input for a new Role record in Nautobot.
	NewRole struct {
		Name         string         `json:"name"`
		ContentTypes []string       `json:"content_types"`
		Color        string         `json:"color,omitempty"`
		Description  string         `json:"description,omitempty"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
		Weight       int            `json:"weight,omitempty"`
	}
)

// RoleGet : Get a Role by UUID identifier.
func (c *Client) RoleGet(id uuid.UUID) (*Role, error) {
	return core.Get[Role](c.Client, extrasEndpointRole, id)
}

// RoleFilter : Get a list of Roles based on query parameters.
func (c *Client) RoleFilter(q *url.Values) ([]Role, error) {
	roles := make([]Role, 0)
	return roles, core.Paginate[Role](c.Client, "extras/roles/", q, &roles)
}

// RoleAll : Get all Roles in Nautobot.
func (c *Client) RoleAll() ([]Role, error) {
	roles := make([]Role, 0)
	return roles, core.Paginate[Role](c.Client, "extras/roles/", nil, &roles)
}

// RoleCreate : Generate a new Role record in Nautobot.
func (c *Client) RoleCreate(obj NewRole) (*Role, error) {
	return core.Create[Role, NewRole](c.Client, extrasEndpointRole, obj)
}

// RoleDelete : Delete a Role by UUID identifier.
func (c *Client) RoleDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointRole, id)
}

// RoleUpdate : Update an existing Role record in Nautobot.
func (c *Client) RoleUpdate(id uuid.UUID, patch map[string]any) (*Role, error) {
	return core.Update[Role](c.Client, extrasEndpointRole, id, patch)
}
