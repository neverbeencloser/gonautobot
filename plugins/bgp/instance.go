package bgp

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

// BGPRoutingInstanceGet : Go function to process requests for the endpoint: /api/plugins/bgp/autonomous-systems/:id/
//
// https://demo.nautobot.com/api/docs/#/plugins/plugins_bgp_autonomous_systems_list
func (c *Client) BGPRoutingInstanceGet(uuid string) (*types.RoutingInstance, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("plugins/bgp/routing-instances/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(types.RoutingInstance)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// BGPRoutingInstanceFilter : Go function to process requests for the endpoint: /api/plugins/bgp/autonomous-systems/
//
// https://demo.nautobot.com/api/docs/#/plugins/plugins_bgp_autonomous_systems_retrieve
func (c *Client) BGPRoutingInstanceFilter(q *url.Values) ([]types.RoutingInstance, error) {
	resp := make([]types.RoutingInstance, 0)
	return resp, core.Paginate[types.RoutingInstance](c.Client, "plugins/bgp/routing-instances/", q, &resp)
}
