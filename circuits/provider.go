package circuits

import (
	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	circuitsEndpointProvider = "circuits/providers/"
)

// ProviderGet : Get a Provider by UUID identifier.
func (c *Client) ProviderGet(id uuid.UUID, opts ...query.Option) (*types.Provider, error) {
	return core.Get[types.Provider](c.Client, circuitsEndpointProvider, id, opts...)
}

// ProviderList : Get a list of Providers in Nautobot, filtered by any query options.
func (c *Client) ProviderList(opts ...query.Option) ([]types.Provider, error) {
	resp := make([]types.Provider, 0)
	return resp, core.Paginate[types.Provider](c.Client, circuitsEndpointProvider, opts, &resp)
}
