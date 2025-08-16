package bgp

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

// BGPAutonomousSystemGet : Go function to process requests for the endpoint: /api/plugins/bgp/autonomous-systems/:id/
//
// https://demo.nautobot.com/api/docs/#/plugins/plugins_bgp_autonomous_systems_list
func (c *Client) BGPAutonomousSystemGet(uuid string) (*types.AutonomousSystem, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("plugins/bgp/autonomous-systems/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(types.AutonomousSystem)
	return ret, c.UnmarshalDo(req, ret)
}

// BGPAutonomousSystemFilter : Go function to process requests for the endpoint: /api/plugins/bgp/autonomous-systems/
//
// https://demo.nautobot.com/api/docs/#/plugins/plugins_bgp_autonomous_systems_retrieve
func (c *Client) BGPAutonomousSystemFilter(q *url.Values) ([]types.AutonomousSystem, error) {
	resp := make([]types.AutonomousSystem, 0)
	return resp, core.Paginate[types.AutonomousSystem](c.Client, "plugins/bgp/autonomous-systems/", q, &resp)
}
