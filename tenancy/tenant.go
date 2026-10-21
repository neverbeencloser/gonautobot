package tenancy

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	tenancyEndpointTenant = "tenancy/tenants/"
)

// TenantList : Get a list of Tenants in Nautobot, filtered by any query options.
//
// https://demo.nautobot.com/api/docs/#/tenancy/tenancy_tenants_list
func (c *Client) TenantList(opts ...query.Option) ([]types.Tenant, error) {
	resp := make([]types.Tenant, 0)
	return resp, core.Paginate[types.Tenant](c.Client, tenancyEndpointTenant, opts, &resp)
}

// TenantGet : Get a Tenant by UUID identifier.
func (c *Client) TenantGet(id uuid.UUID, opts ...query.Option) (*types.Tenant, error) {
	return core.Get[types.Tenant](c.Client, tenancyEndpointTenant, id, opts...)
}

// TenantCreate : Generate a new Tenant record in Nautobot.
func (c *Client) TenantCreate(obj types.NewTenant) (*types.Tenant, error) {
	return core.Create[types.Tenant, types.NewTenant](c.Client, tenancyEndpointTenant, obj)
}

// TenantDelete : Delete a Tenant by UUID identifier.
func (c *Client) TenantDelete(id uuid.UUID) error {
	return core.Delete(c.Client, tenancyEndpointTenant, id)
}

// TenantUpdate : Update an existing Tenant record in Nautobot.
func (c *Client) TenantUpdate(id uuid.UUID, patch map[string]any) (*types.Tenant, error) {
	return core.Update[types.Tenant](c.Client, tenancyEndpointTenant, id, patch)
}
