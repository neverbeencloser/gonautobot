package dcim

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointSoftwareVersion = "dcim/software-versions/"
)

// SoftwareVersionGet : Get a SoftwareVersion by UUID identifier.
func (c *Client) SoftwareVersionGet(id uuid.UUID) (*types.SoftwareVersion, error) {
	return core.Get[types.SoftwareVersion](c.Client, dcimEndpointSoftwareVersion, id)
}

// SoftwareVersionFilter : Get a list of SoftwareVersions based on query parameters.
func (c *Client) SoftwareVersionFilter(q *url.Values) ([]types.SoftwareVersion, error) {
	softwareVersions := make([]types.SoftwareVersion, 0)
	return softwareVersions, core.Paginate[types.SoftwareVersion](c.Client, dcimEndpointSoftwareVersion, q, &softwareVersions)
}

// SoftwareVersionAll : Get all SoftwareVersions in Nautobot.
func (c *Client) SoftwareVersionAll() ([]types.SoftwareVersion, error) {
	softwareVersions := make([]types.SoftwareVersion, 0)
	return softwareVersions, core.Paginate[types.SoftwareVersion](c.Client, dcimEndpointSoftwareVersion, nil, &softwareVersions)
}

// SoftwareVersionCreate : Generate a new SoftwareVersion record in Nautobot.
func (c *Client) SoftwareVersionCreate(obj types.NewSoftwareVersion) (*types.SoftwareVersion, error) {
	return core.Create[types.SoftwareVersion, types.NewSoftwareVersion](c.Client, dcimEndpointSoftwareVersion, obj)
}

// SoftwareVersionDelete : Delete a SoftwareVersion by UUID identifier.
func (c *Client) SoftwareVersionDelete(id uuid.UUID) error {
	return core.Delete(c.Client, dcimEndpointSoftwareVersion, id)
}

// SoftwareVersionUpdate : Update an existing SoftwareVersion record in Nautobot.
func (c *Client) SoftwareVersionUpdate(id uuid.UUID, patch map[string]any) (*types.SoftwareVersion, error) {
	return core.Update[types.SoftwareVersion](c.Client, dcimEndpointSoftwareVersion, id, patch)
}
