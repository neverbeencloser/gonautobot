package dcim

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointSoftwareImageFile = "dcim/software-image-files/"
)

// SoftwareImageFileGet : Get a SoftwareImageFile by UUID identifier.
func (c *Client) SoftwareImageFileGet(id uuid.UUID) (*types.SoftwareImageFile, error) {
	return core.Get[types.SoftwareImageFile](c.Client, dcimEndpointSoftwareImageFile, id)
}

// SoftwareImageFileFilter : Get a list of SoftwareImageFiles based on query parameters.
func (c *Client) SoftwareImageFileFilter(q *url.Values) ([]types.SoftwareImageFile, error) {
	sifs := make([]types.SoftwareImageFile, 0)
	return sifs, core.Paginate[types.SoftwareImageFile](c.Client, dcimEndpointSoftwareImageFile, q, &sifs)
}

// SoftwareImageFileAll : Get all SoftwareImageFiles in Nautobot.
func (c *Client) SoftwareImageFileAll() ([]types.SoftwareImageFile, error) {
	sifs := make([]types.SoftwareImageFile, 0)
	return sifs, core.Paginate[types.SoftwareImageFile](c.Client, dcimEndpointSoftwareImageFile, nil, &sifs)
}

// SoftwareImageFileCreate : Generate a new SoftwareImageFile record in Nautobot.
func (c *Client) SoftwareImageFileCreate(obj types.NewSoftwareImageFile) (*types.SoftwareImageFile, error) {
	return core.Create[types.SoftwareImageFile, types.NewSoftwareImageFile](c.Client, dcimEndpointSoftwareImageFile, obj)
}

// SoftwareImageFileDelete : Delete a SoftwareImageFile by UUID identifier.
func (c *Client) SoftwareImageFileDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointSoftwareImageFile, id)
}

// SoftwareImageFileUpdate : Update an existing SoftwareImageFile record in Nautobot.
func (c *Client) SoftwareImageFileUpdate(id uuid.UUID, patch map[string]any) (*types.SoftwareImageFile, error) {
	return core.Update[types.SoftwareImageFile](c.Client, dcimEndpointSoftwareImageFile, id, patch)
}
