package circuits

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	circuitsEndpointProvider = "circuits/providers/"
)

// ProviderGet : Get a Provider by UUID identifier.
func (c *Client) ProviderGet(id uuid.UUID) (*types.Provider, error) {
	return core.Get[types.Provider](c.Client, circuitsEndpointProvider, id)
}

// ProviderFilter : Get a list of Providers based on query parameters.
func (c *Client) ProviderFilter(q *url.Values) ([]types.Provider, error) {
	resp := make([]types.Provider, 0)
	return resp, core.Paginate[types.Provider](c.Client, circuitsEndpointProvider, q, &resp)
}

// ProviderAll : Get all Providers in Nautobot.
func (c *Client) ProviderAll() ([]types.Provider, error) {
	resp := make([]types.Provider, 0)
	return resp, core.Paginate[types.Provider](c.Client, circuitsEndpointProvider, nil, &resp)
}
