package dcim

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointInterface = "dcim/interfaces/"
)

// InterfaceGet : Go function to process requests for the endpoint: /api/dcim/interfaces/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_interfaces_retrieve
func (c *Client) InterfaceGet(id uuid.UUID, opts ...query.Option) (*types.DeviceInterface, error) {
	return core.Get[types.DeviceInterface](c.Client, dcimEndpointInterface, id, opts...)
}

// InterfaceList : Get a list of Interfaces in Nautobot, filtered by any query options.
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_interfaces_list
func (c *Client) InterfaceList(opts ...query.Option) ([]types.DeviceInterface, error) {
	resp := make([]types.DeviceInterface, 0)
	return resp, core.Paginate[types.DeviceInterface](c.Client, dcimEndpointInterface, opts, &resp)
}
