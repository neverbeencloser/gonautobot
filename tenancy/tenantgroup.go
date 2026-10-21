package tenancy

import (
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

// TenantGroupList : Get a list of Tenant Groups in Nautobot, filtered by any query options.
//
// https://demo.nautobot.com/api/docs/#/tenancy/tenancy_tenant_groups_list
func (c *Client) TenantGroupList(opts ...query.Option) ([]types.TenantGroup, error) {
	resp := make([]types.TenantGroup, 0)
	return resp, core.Paginate[types.TenantGroup](c.Client, "tenancy/tenant-groups/", opts, &resp)
}
