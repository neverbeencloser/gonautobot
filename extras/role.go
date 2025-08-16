package extras

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	extrasEndpointRole = "extras/roles/"
)

// RoleGet : Get a Role by UUID identifier.
func (c *Client) RoleGet(id uuid.UUID) (*types.Role, error) {
	return core.Get[types.Role](c.Client, extrasEndpointRole, id)
}

// RoleFilter : Get a list of Roles based on query parameters.
func (c *Client) RoleFilter(q *url.Values) ([]types.Role, error) {
	roles := make([]types.Role, 0)
	return roles, core.Paginate[types.Role](c.Client, "extras/roles/", q, &roles)
}

// RoleAll : Get all Roles in Nautobot.
func (c *Client) RoleAll() ([]types.Role, error) {
	roles := make([]types.Role, 0)
	return roles, core.Paginate[types.Role](c.Client, "extras/roles/", nil, &roles)
}

// RoleCreate : Generate a new Role record in Nautobot.
func (c *Client) RoleCreate(obj types.NewRole) (*types.Role, error) {
	return core.Create[types.Role, types.NewRole](c.Client, extrasEndpointRole, obj)
}

// RoleDelete : Delete a Role by UUID identifier.
func (c *Client) RoleDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointRole, id)
}

// RoleUpdate : Update an existing Role record in Nautobot.
func (c *Client) RoleUpdate(id uuid.UUID, patch map[string]any) (*types.Role, error) {
	return core.Update[types.Role](c.Client, extrasEndpointRole, id, patch)
}
