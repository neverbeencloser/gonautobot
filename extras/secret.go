package extras

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	extrasEndpointSecret = "extras/secrets/"
)

// SecretGet : Get a Secret by UUID identifier.
func (c *Client) SecretGet(id uuid.UUID) (*types.Secret, error) {
	return core.Get[types.Secret](c.Client, extrasEndpointSecret, id)
}

// SecretFilter : Get a list of Secrets based on query parameters.
func (c *Client) SecretFilter(q *url.Values) ([]types.Secret, error) {
	secrets := make([]types.Secret, 0)
	return secrets, core.Paginate[types.Secret](c.Client, extrasEndpointSecret, q, &secrets)
}

// SecretAll : Get all Secrets in Nautobot.
func (c *Client) SecretAll() ([]types.Secret, error) {
	secrets := make([]types.Secret, 0)
	return secrets, core.Paginate[types.Secret](c.Client, extrasEndpointSecret, nil, &secrets)
}

// SecretCreate : Generate a new Secret record in Nautobot.
func (c *Client) SecretCreate(obj types.NewSecret) (*types.Secret, error) {
	return core.Create[types.Secret, types.NewSecret](c.Client, extrasEndpointSecret, obj)
}

// SecretDelete : Delete a Secret by UUID identifier.
func (c *Client) SecretDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointSecret, id)
}

// SecretUpdate : Update an existing Secret record in Nautobot.
func (c *Client) SecretUpdate(id uuid.UUID, patch map[string]any) (*types.Secret, error) {
	return core.Update[types.Secret](c.Client, extrasEndpointSecret, id, patch)
}
