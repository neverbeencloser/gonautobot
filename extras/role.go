package extras

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	extrasEndpointRole = "extras/roles/"
)

// RoleGet : Get a Role by UUID identifier.
func (c *Client) RoleGet(id uuid.UUID, opts ...query.Option) (*types.Role, error) {
	return core.Get[types.Role](c.Client, extrasEndpointRole, id, opts...)
}

// RoleList : Get a list of Roles in Nautobot, filtered by any query options.
func (c *Client) RoleList(opts ...query.Option) ([]types.Role, error) {
	roles := make([]types.Role, 0)
	return roles, core.Paginate[types.Role](c.Client, "extras/roles/", opts, &roles)
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
