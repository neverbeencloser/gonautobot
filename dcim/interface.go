package dcim

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

// InterfaceGet : Go function to process requests for the endpoint: /api/dcim/interfaces/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_interfaces_retrieve
func (c *Client) InterfaceGet(uuid string) (*types.DeviceInterface, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/interfaces/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(types.DeviceInterface)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// InterfaceFilter : Go function to process requests for the endpoint: /api/dcim/interfaces/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_interfaces_list
func (c *Client) InterfaceFilter(q *url.Values) ([]types.DeviceInterface, error) {
	resp := make([]types.DeviceInterface, 0)
	if err := core.Paginate[types.DeviceInterface](c.Client, "dcim/interfaces/", q, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
