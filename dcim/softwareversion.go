package dcim

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	dcimEndpointSoftwareVersion = "dcim/software-versions/"
)

// SoftwareVersionGet : Get a SoftwareVersion by UUID identifier.
func (c *Client) SoftwareVersionGet(id uuid.UUID, opts ...query.Option) (*types.SoftwareVersion, error) {
	return core.Get[types.SoftwareVersion](c.Client, dcimEndpointSoftwareVersion, id, opts...)
}

// SoftwareVersionList : Get a list of SoftwareVersions in Nautobot, filtered by any query options.
func (c *Client) SoftwareVersionList(opts ...query.Option) ([]types.SoftwareVersion, error) {
	softwareVersions := make([]types.SoftwareVersion, 0)
	return softwareVersions, core.Paginate[types.SoftwareVersion](c.Client, dcimEndpointSoftwareVersion, opts, &softwareVersions)
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
