package dcim

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/extras"
	"github.com/neverbeencloser/gonautobot/shared"
)

const (
	dcimEndpointDeviceType = "dcim/device-types/"
)

type (
	// DeviceType represents a device type in Nautobot.
	DeviceType struct {
		ID                 uuid.UUID           `json:"id"`
		Comments           string              `json:"comments"`
		Created            time.Time           `json:"created"`
		CustomFields       map[string]any      `json:"custom_fields"`
		DeviceCount        int                 `json:"device_count"`
		DeviceFamily       *DeviceFamily       `json:"device_family"`
		Display            string              `json:"display"`
		FrontImage         *string             `json:"front_image"`
		IsFullDepth        bool                `json:"is_full_depth"`
		LastUpdated        time.Time           `json:"last_updated"`
		Manufacturer       *Manufacturer       `json:"manufacturer"`
		Model              string              `json:"model"`
		NaturalSlug        string              `json:"natural_slug"`
		NotesURL           string              `json:"notes_url"`
		ObjectType         string              `json:"object_type"`
		PartNumber         string              `json:"part_number"`
		RearImage          *string             `json:"rear_image"`
		SoftwareImageFiles []SoftwareImageFile `json:"software_image_files"`
		SubDeviceRole      *shared.LabelValue  `json:"subdevice_role"`
		Tags               []extras.Tag        `json:"tags"`
		UHeight            int                 `json:"u_height"`
		URL                string              `json:"url"`
	}

	// NewDeviceType represents the structure for creating a new DeviceType in Nautobot.
	//
	// It uses our internal 'form' tags for generating multipart/form-data requests.
	NewDeviceType struct {
		Manufacturer       string         `form:"manufacturer"`
		Model              string         `form:"model"`
		Comments           string         `form:"comments"`
		CustomFields       map[string]any `form:"custom_fields,omitempty"`
		DeviceFamily       string         `form:"device_family"`
		FrontImage         *string        `form:"front_image,omitempty,upload"`
		IsFullDepth        bool           `form:"is_full_depth"`
		PartNumber         string         `form:"part_number"`
		RearImage          *string        `form:"rear_image,omitempty,upload"`
		SoftwareImageFiles []string       `form:"software_image_files,omitempty"`
		SubDeviceRole      string         `form:"subdevice_role"`
		Tags               []string       `form:"tags,omitempty"`
		UHeight            int            `form:"u_height"`
	}
)

// DeviceTypeGet : Get a DeviceType by UUID identifier.
func (c *Client) DeviceTypeGet(id uuid.UUID) (*DeviceType, error) {
	return core.Get[DeviceType](c.Client, dcimEndpointDeviceType, id)
}

// DeviceTypeFilter : Get a list of DeviceTypes based on query parameters.
func (c *Client) DeviceTypeFilter(q *url.Values) ([]DeviceType, error) {
	devicetypes := make([]DeviceType, 0)
	return devicetypes, core.Paginate[DeviceType](c.Client, dcimEndpointDeviceType, q, &devicetypes)
}

// DeviceTypeAll : Get all DeviceTypes in Nautobot.
func (c *Client) DeviceTypeAll() ([]DeviceType, error) {
	devicetypes := make([]DeviceType, 0)
	return devicetypes, core.Paginate[DeviceType](c.Client, dcimEndpointDeviceType, nil, &devicetypes)
}

// DeviceTypeCreate : Generate a new DeviceType record in Nautobot.
func (c *Client) DeviceTypeCreate(obj NewDeviceType) (*DeviceType, error) {
	return core.CreateMultipart[DeviceType, NewDeviceType](c.Client, dcimEndpointDeviceType, obj)
}

// DeviceTypeDelete : Delete a DeviceType by UUID identifier.
func (c *Client) DeviceTypeDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointDeviceType, id)
}

// DeviceTypeUpdate : Update an existing DeviceType record in Nautobot.
func (c *Client) DeviceTypeUpdate(id uuid.UUID, obj NewDeviceType) (*DeviceType, error) {
	return core.UpdateMultipart[DeviceType, NewDeviceType](c.Client, dcimEndpointDeviceType, id, obj)
}

// SetFrontImage is a convenience method to set the front image path for a NewDeviceType.
func (n *NewDeviceType) SetFrontImage(path string) {
	n.FrontImage = &path
}

// SetRearImage is a convenience method to set the rear image path for a NewDeviceType.
func (n *NewDeviceType) SetRearImage(path string) {
	n.RearImage = &path
}
