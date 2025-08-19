package circuits

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	circuitsEndpointCircuit = "circuits/circuits/"
)

// CircuitGet : Get a Circuit by UUID identifier.
func (c *Client) CircuitGet(id uuid.UUID) (*types.Circuit, error) {
	return core.Get[types.Circuit](c.Client, circuitsEndpointCircuit, id)
}

// CircuitFilter : Get a list of Circuits based on query parameters.
func (c *Client) CircuitFilter(q *url.Values) ([]types.Circuit, error) {
	resp := make([]types.Circuit, 0)
	return resp, core.Paginate[types.Circuit](c.Client, circuitsEndpointCircuit, q, &resp)
}

// CircuitAll : Get all Circuits in Nautobot.
func (c *Client) CircuitAll() ([]types.Circuit, error) {
	resp := make([]types.Circuit, 0)
	return resp, core.Paginate[types.Circuit](c.Client, circuitsEndpointCircuit, nil, &resp)
}

// CircuitDelete : Delete a Circuit by UUID identifier.
func (c *Client) CircuitDelete(id uuid.UUID) error {
	return core.Delete(c.Client, circuitsEndpointCircuit, id)
}

// CircuitCreate : Create a new Circuit record in Nautobot.
func (c *Client) CircuitCreate(obj types.CircuitRequest) (*types.Circuit, error) {
	return core.Create[types.Circuit, types.CircuitRequest](c.Client, circuitsEndpointCircuit, obj)
}

// CircuitUpdate : Update an existing Circuit record in Nautobot.
func (c *Client) CircuitUpdate(id uuid.UUID, patch map[string]any) (*types.Circuit, error) {
	return core.Update[types.Circuit](c.Client, circuitsEndpointCircuit, id, patch)
}
