package bgp

import (
	"encoding/json"
	"fmt"
	nautobot "github.com/josh-silvas/gonautobot/extras"
	nautobot2 "github.com/josh-silvas/gonautobot/shared"
	"github.com/josh-silvas/gonautobot/shared/nested"
	"net/http"
	"net/url"
)

type (
	// RoutingInstance : Routing Instance data representation in Nautobot.
	RoutingInstance struct {
		AutonomousSystem *AutonomousSystem        `json:"autonomous_system"`
		Created          string                   `json:"created"`
		CustomFields     map[string]interface{}   `json:"custom_fields"`
		Description      string                   `json:"description"`
		Device           *nested.Device           `json:"device"`
		Display          string                   `json:"display"`
		Endpoints        []nested.BGPPeerEndpoint `json:"endpoints"`
		ExtraAttributes  map[string]interface{}   `json:"extra_attributes"`
		ID               string                   `json:"id"`
		LastUpdated      string                   `json:"last_updated"`
		RouterID         *nested.IPAddress        `json:"router_id"`
		Status           *nautobot2.LabelValue    `json:"status"`
		Tags             []nautobot.Tag           `json:"tags"`
		URL              string                   `json:"url"`
	}
)

// BGPRoutingInstance : Go function to process requests for the endpoint: /api/plugins/bgp/autonomous-systems/:id/
//
// https://demo.nautobot.com/api/docs/#/plugins/plugins_bgp_autonomous_systems_list
func (c *Client) BGPRoutingInstance(uuid string, q *url.Values) (*RoutingInstance, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("plugins/bgp/routing-instances/%s/", url.PathEscape(uuid)), nil, q)
	if err != nil {
		return nil, err
	}

	ret := new(RoutingInstance)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// BGPRoutingInstances : Go function to process requests for the endpoint: /api/plugins/bgp/autonomous-systems/
//
// https://demo.nautobot.com/api/docs/#/plugins/plugins_bgp_autonomous_systems_retrieve
func (c *Client) BGPRoutingInstances(q *url.Values) ([]RoutingInstance, error) {
	req, err := c.Request(http.MethodGet, "plugins/bgp/routing-instances/", nil, q)
	if err != nil {
		return nil, err
	}

	resp := new(nautobot2.ResponseList)
	ret := make([]RoutingInstance, 0)

	if err = c.UnmarshalDo(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("BGPRoutingInstances.error.json.Unmarshal(%w)", err)
	}
	return ret, err
}
