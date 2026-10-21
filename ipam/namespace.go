package ipam

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	ipamEndpointNamespace = "ipam/namespaces/"
)

// NamespaceGet : Get a Namespace by UUID identifier.
func (c *Client) NamespaceGet(id uuid.UUID, opts ...query.Option) (*types.Namespace, error) {
	return core.Get[types.Namespace](c.Client, ipamEndpointNamespace, id, opts...)
}

// NamespaceList : Get a list of Namespaces in Nautobot, filtered by any query options.
func (c *Client) NamespaceList(opts ...query.Option) ([]types.Namespace, error) {
	namespaces := make([]types.Namespace, 0)
	return namespaces, core.Paginate[types.Namespace](c.Client, ipamEndpointNamespace, opts, &namespaces)
}

// NamespaceCreate : Generate a new Namespace record in Nautobot.
func (c *Client) NamespaceCreate(obj types.NewNamespace) (*types.Namespace, error) {
	return core.Create[types.Namespace, types.NewNamespace](c.Client, ipamEndpointNamespace, obj)
}

// NamespaceUpdate : Update an existing Namespace record in Nautobot.
func (c *Client) NamespaceUpdate(id uuid.UUID, patch map[string]any) (*types.Namespace, error) {
	return core.Update[types.Namespace](c.Client, ipamEndpointNamespace, id, patch)
}

// NamespaceDelete : Delete a Namespace by UUID identifier.
func (c *Client) NamespaceDelete(id uuid.UUID) error {
	return core.Delete(c.Client, ipamEndpointNamespace, id)
}
