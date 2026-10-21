package dcim

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointManufacturer = "dcim/manufacturers/"
)

// ManufacturerGet : Get a Manufacturer by UUID identifier.
func (c *Client) ManufacturerGet(id uuid.UUID, opts ...query.Option) (*types.Manufacturer, error) {
	return core.Get[types.Manufacturer](c.Client, dcimEndpointManufacturer, id, opts...)
}

// ManufacturerList : Get a list of Manufacturers in Nautobot, filtered by any query options.
func (c *Client) ManufacturerList(opts ...query.Option) ([]types.Manufacturer, error) {
	manufacturers := make([]types.Manufacturer, 0)
	return manufacturers, core.Paginate[types.Manufacturer](c.Client, dcimEndpointManufacturer, opts, &manufacturers)
}

// ManufacturerCreate : Generate a new Manufacturer record in Nautobot.
func (c *Client) ManufacturerCreate(obj types.NewManufacturer) (*types.Manufacturer, error) {
	return core.Create[types.Manufacturer, types.NewManufacturer](c.Client, dcimEndpointManufacturer, obj)
}

// ManufacturerDelete : Delete a Manufacturer by UUID identifier.
func (c *Client) ManufacturerDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointManufacturer, id)
}

// ManufacturerUpdate : Update an existing Manufacturer record in Nautobot.
func (c *Client) ManufacturerUpdate(id uuid.UUID, patch map[string]any) (*types.Manufacturer, error) {
	return core.Update[types.Manufacturer](c.Client, dcimEndpointManufacturer, id, patch)
}
