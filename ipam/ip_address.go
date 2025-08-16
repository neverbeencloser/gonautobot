package ipam

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

// IPAddressGet : Go function to process requests for the endpoint: /api/ipam/ip_addresses/:id/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_ip_addresses_retrieve
func (c *Client) IPAddressGet(uuid string) (*types.IPAddress, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("ipam/ip-addresses/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(types.IPAddress)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// IPAddressFilter : Go function to process requests for the endpoint: /api/ipam/ip_addresses/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_ip_addresses_list
func (c *Client) IPAddressFilter(q *url.Values) ([]types.IPAddress, error) {
	resp := make([]types.IPAddress, 0)
	return resp, core.Paginate[types.IPAddress](c.Client, "ipam/ip-addresses/", q, &resp)
}
