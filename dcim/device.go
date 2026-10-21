package dcim

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointDevice = "dcim/devices/"
)

// DeviceGet : Get a Device by UUID identifier.
func (c *Client) DeviceGet(id uuid.UUID, opts ...query.Option) (*types.Device, error) {
	return core.Get[types.Device](c.Client, dcimEndpointDevice, id, opts...)
}

// DeviceList : Get a list of Devices in Nautobot, filtered by any query options.
func (c *Client) DeviceList(opts ...query.Option) ([]types.Device, error) {
	devices := make([]types.Device, 0)
	return devices, core.Paginate[types.Device](c.Client, dcimEndpointDevice, opts, &devices)
}

// DeviceCreate : Generate a new Device record in Nautobot.
func (c *Client) DeviceCreate(obj types.NewDevice) (*types.Device, error) {
	return core.Create[types.Device, types.NewDevice](c.Client, dcimEndpointDevice, obj)
}

// DeviceDelete : Delete a Device by UUID identifier.
func (c *Client) DeviceDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointDevice, id)
}

// DeviceUpdate : Update an existing Device record in Nautobot.
func (c *Client) DeviceUpdate(id uuid.UUID, patch map[string]any) (*types.Device, error) {
	return core.Update[types.Device](c.Client, dcimEndpointDevice, id, patch)
}
