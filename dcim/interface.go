package dcim

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointInterface = "dcim/interfaces/"
)

// InterfaceGet : Go function to process requests for the endpoint: /api/dcim/interfaces/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_interfaces_retrieve
func (c *Client) InterfaceGet(id uuid.UUID) (*types.DeviceInterface, error) {
	return core.Get[types.DeviceInterface](c.Client, dcimEndpointInterface, id)
}

// InterfaceFilter : Go function to process requests for the endpoint: /api/dcim/interfaces/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_interfaces_list
func (c *Client) InterfaceFilter(q *url.Values) ([]types.DeviceInterface, error) {
	resp := make([]types.DeviceInterface, 0)
	return resp, core.Paginate[types.DeviceInterface](c.Client, dcimEndpointInterface, q, &resp)
}
