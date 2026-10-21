package dcim

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointDeviceFamily = "dcim/device-families/"
)

// DeviceFamilyGet : Get a DeviceFamily by UUID identifier.
func (c *Client) DeviceFamilyGet(id uuid.UUID, opts ...query.Option) (*types.DeviceFamily, error) {
	return core.Get[types.DeviceFamily](c.Client, dcimEndpointDeviceFamily, id, opts...)
}

// DeviceFamilyList : Get a list of DeviceFamilies in Nautobot, filtered by any query options.
func (c *Client) DeviceFamilyList(opts ...query.Option) ([]types.DeviceFamily, error) {
	deviceFamilies := make([]types.DeviceFamily, 0)
	return deviceFamilies, core.Paginate[types.DeviceFamily](c.Client, dcimEndpointDeviceFamily, opts, &deviceFamilies)
}

// DeviceFamilyCreate : Generate a new DeviceFamily record in Nautobot.
func (c *Client) DeviceFamilyCreate(obj types.NewDeviceFamily) (*types.DeviceFamily, error) {
	return core.Create[types.DeviceFamily, types.NewDeviceFamily](c.Client, dcimEndpointDeviceFamily, obj)
}

// DeviceFamilyDelete : Delete a DeviceFamily by UUID identifier.
func (c *Client) DeviceFamilyDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointDeviceFamily, id)
}

// DeviceFamilyUpdate : Update an existing DeviceFamily record in Nautobot.
func (c *Client) DeviceFamilyUpdate(id uuid.UUID, patch map[string]any) (*types.DeviceFamily, error) {
	return core.Update[types.DeviceFamily](c.Client, dcimEndpointDeviceFamily, id, patch)
}
