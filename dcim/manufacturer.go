package dcim

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/josh-silvas/gonautobot/core"
)

const (
	dcimEndpointManufacturer = "dcim/manufacturers/"
)

type (
	// Manufacturer represents a manufacturer in Nautobot's DCIM application.
	Manufacturer struct {
		ID                 uuid.UUID      `json:"id"`
		CloudAccountCount  int            `json:"cloud_account_count"`
		Created            time.Time      `json:"created"`
		CustomFields       map[string]any `json:"custom_fields"`
		Description        string         `json:"description"`
		DeviceTypeCount    int            `json:"device_type_count"`
		Display            string         `json:"display"`
		InventoryItemCount int            `json:"inventory_item_count"`
		LastUpdated        time.Time      `json:"last_updated"`
		Name               string         `json:"name"`
		NaturalSlug        string         `json:"natural_slug"`
		NotesURL           string         `json:"notes_url"`
		ObjectType         string         `json:"object_type"`
		PlatformCount      int            `json:"platform_count"`
		URL                string         `json:"url"`
	}

	// NewManufacturer represents the data required to create a new manufacturer.
	NewManufacturer struct {
		Name         string         `json:"name"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
		Description  string         `json:"description,omitempty"`
	}
)

// ManufacturerGet : Get a Manufacturer by UUID identifier.
func (c *Client) ManufacturerGet(id uuid.UUID) (*Manufacturer, error) {
	return core.Get[Manufacturer](c.Client, dcimEndpointManufacturer, id)
}

// ManufacturerFilter : Get a list of Manufacturers based on query parameters.
func (c *Client) ManufacturerFilter(q *url.Values) ([]Manufacturer, error) {
	manufacturers := make([]Manufacturer, 0)
	return manufacturers, core.Paginate[Manufacturer](c.Client, dcimEndpointManufacturer, q, &manufacturers)
}

// ManufacturerAll : Get all Manufacturers in Nautobot.
func (c *Client) ManufacturerAll() ([]Manufacturer, error) {
	manufacturers := make([]Manufacturer, 0)
	return manufacturers, core.Paginate[Manufacturer](c.Client, dcimEndpointManufacturer, nil, &manufacturers)
}

// ManufacturerCreate : Generate a new Manufacturer record in Nautobot.
func (c *Client) ManufacturerCreate(obj NewManufacturer) (*Manufacturer, error) {
	return core.Create[Manufacturer, NewManufacturer](c.Client, dcimEndpointManufacturer, obj)
}

// ManufacturerDelete : Delete a Manufacturer by UUID identifier.
func (c *Client) ManufacturerDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointManufacturer, id)
}

// ManufacturerUpdate : Update an existing Manufacturer record in Nautobot.
func (c *Client) ManufacturerUpdate(id uuid.UUID, patch map[string]any) (*Manufacturer, error) {
	return core.Update[Manufacturer](c.Client, dcimEndpointManufacturer, id, patch)
}
