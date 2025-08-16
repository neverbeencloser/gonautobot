package tenancy

import (
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

// TenantGroupFilter : Go function to process requests for the endpoint: /api/tenancy/tenant-groups/
//
// https://demo.nautobot.com/api/docs/#/tenancy/tenancy_tenant_groups_list
func (c *Client) TenantGroupFilter(q *url.Values) ([]types.TenantGroup, error) {
	resp := make([]types.TenantGroup, 0)
	return resp, core.Paginate[types.TenantGroup](c.Client, "tenancy/tenant-groups/", q, &resp)
}
