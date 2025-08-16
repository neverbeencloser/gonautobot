package ipam

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

// VLANGet : Go function to process requests for the endpoint: /api/ipam/vlans/:id/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_vlans_retrieve
func (c *Client) VLANGet(uuid string) (*types.VLAN, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("ipam/vlans/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}
	ret := new(types.VLAN)
	return ret, c.UnmarshalDo(req, ret)
}

// VLANFilter : Go function to process requests for the endpoint: /api/ipam/vlans/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_vlans_list
func (c *Client) VLANFilter(q *url.Values) ([]types.VLAN, error) {
	resp := make([]types.VLAN, 0)
	return resp, core.Paginate[types.VLAN](c.Client, "ipam/vlans/", q, &resp)
}
