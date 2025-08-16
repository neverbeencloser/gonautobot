package dcim

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointDeviceFamily = "dcim/device-families/"
)

// DeviceFamilyGet : Get a DeviceFamily by UUID identifier.
func (c *Client) DeviceFamilyGet(id uuid.UUID) (*types.DeviceFamily, error) {
	return core.Get[types.DeviceFamily](c.Client, dcimEndpointDeviceFamily, id)
}

// DeviceFamilyFilter : Get a list of DeviceFamilies based on query parameters.
func (c *Client) DeviceFamilyFilter(q *url.Values) ([]types.DeviceFamily, error) {
	deviceFamilies := make([]types.DeviceFamily, 0)
	return deviceFamilies, core.Paginate[types.DeviceFamily](c.Client, dcimEndpointDeviceFamily, q, &deviceFamilies)
}

// DeviceFamilyAll : Get all DeviceFamilies in Nautobot.
func (c *Client) DeviceFamilyAll() ([]types.DeviceFamily, error) {
	deviceFamilies := make([]types.DeviceFamily, 0)
	return deviceFamilies, core.Paginate[types.DeviceFamily](c.Client, dcimEndpointDeviceFamily, nil, &deviceFamilies)
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
