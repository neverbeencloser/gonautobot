package dcim

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointDeviceFamily = "dcim/device-families/"
)

type (
	// DeviceFamily represents a device family in Nautobot's DCIM module.
	DeviceFamily struct {
		ID           uuid.UUID      `json:"id"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		Tags         []types.Tag    `json:"tags"`
		URL          string         `json:"url"`
	}

	// NewDeviceFamily represents the data required to create a new device family.
	NewDeviceFamily struct {
		Name         string         `json:"name"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
		Description  string         `json:"description,omitempty"`
		Tags         []string       `json:"tags,omitempty"`
	}
)

// DeviceFamilyGet : Get a DeviceFamily by UUID identifier.
func (c *Client) DeviceFamilyGet(id uuid.UUID) (*DeviceFamily, error) {
	return core.Get[DeviceFamily](c.Client, dcimEndpointDeviceFamily, id)
}

// DeviceFamilyFilter : Get a list of DeviceFamilies based on query parameters.
func (c *Client) DeviceFamilyFilter(q *url.Values) ([]DeviceFamily, error) {
	deviceFamilies := make([]DeviceFamily, 0)
	return deviceFamilies, core.Paginate[DeviceFamily](c.Client, dcimEndpointDeviceFamily, q, &deviceFamilies)
}

// DeviceFamilyAll : Get all DeviceFamilies in Nautobot.
func (c *Client) DeviceFamilyAll() ([]DeviceFamily, error) {
	deviceFamilies := make([]DeviceFamily, 0)
	return deviceFamilies, core.Paginate[DeviceFamily](c.Client, dcimEndpointDeviceFamily, nil, &deviceFamilies)
}

// DeviceFamilyCreate : Generate a new DeviceFamily record in Nautobot.
func (c *Client) DeviceFamilyCreate(obj NewDeviceFamily) (*DeviceFamily, error) {
	return core.Create[DeviceFamily, NewDeviceFamily](c.Client, dcimEndpointDeviceFamily, obj)
}

// DeviceFamilyDelete : Delete a DeviceFamily by UUID identifier.
func (c *Client) DeviceFamilyDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointDeviceFamily, id)
}

// DeviceFamilyUpdate : Update an existing DeviceFamily record in Nautobot.
func (c *Client) DeviceFamilyUpdate(id uuid.UUID, patch map[string]any) (*DeviceFamily, error) {
	return core.Update[DeviceFamily](c.Client, dcimEndpointDeviceFamily, id, patch)
}
