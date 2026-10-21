package dcim

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointDeviceType = "dcim/device-types/"
)

// DeviceTypeGet : Get a DeviceType by UUID identifier.
func (c *Client) DeviceTypeGet(id uuid.UUID, opts ...query.Option) (*types.DeviceType, error) {
	return core.Get[types.DeviceType](c.Client, dcimEndpointDeviceType, id, opts...)
}

// DeviceTypeList : Get a list of DeviceTypes in Nautobot, filtered by any query options.
func (c *Client) DeviceTypeList(opts ...query.Option) ([]types.DeviceType, error) {
	devicetypes := make([]types.DeviceType, 0)
	return devicetypes, core.Paginate[types.DeviceType](c.Client, dcimEndpointDeviceType, opts, &devicetypes)
}

// DeviceTypeCreate : Generate a new DeviceType record in Nautobot.
func (c *Client) DeviceTypeCreate(obj types.NewDeviceType) (*types.DeviceType, error) {
	return core.CreateMultipart[types.DeviceType, types.NewDeviceType](c.Client, dcimEndpointDeviceType, obj)
}

// DeviceTypeDelete : Delete a DeviceType by UUID identifier.
func (c *Client) DeviceTypeDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointDeviceType, id)
}

// DeviceTypeUpdate : Update an existing DeviceType record in Nautobot.
func (c *Client) DeviceTypeUpdate(id uuid.UUID, obj types.NewDeviceType) (*types.DeviceType, error) {
	return core.UpdateMultipart[types.DeviceType, types.NewDeviceType](c.Client, dcimEndpointDeviceType, id, obj)
}
