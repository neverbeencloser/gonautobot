package extras

import (
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	extrasEndpointCustomField = "extras/custom-fields/"
)

// CustomFieldGet : Get a CustomField by UUID identifier.
func (c *Client) CustomFieldGet(id uuid.UUID) (*types.CustomField, error) {
	return core.Get[types.CustomField](c.Client, extrasEndpointCustomField, id)
}

// CustomFieldFilter : Get a list of CustomFields based on query parameters.
func (c *Client) CustomFieldFilter(q *url.Values) ([]types.CustomField, error) {
	customFields := make([]types.CustomField, 0)
	return customFields, core.Paginate[types.CustomField](c.Client, extrasEndpointCustomField, q, &customFields)
}

// CustomFieldAll : Get all CustomFields in Nautobot.
func (c *Client) CustomFieldAll() ([]types.CustomField, error) {
	customFields := make([]types.CustomField, 0)
	return customFields, core.Paginate[types.CustomField](c.Client, extrasEndpointCustomField, nil, &customFields)
}

// CustomFieldCreate : Generate a new CustomField record in Nautobot.
func (c *Client) CustomFieldCreate(obj types.NewCustomField) (*types.CustomField, error) {
	return core.Create[types.CustomField, types.NewCustomField](c.Client, extrasEndpointCustomField, obj)
}

// CustomFieldDelete : Delete a CustomField by UUID identifier.
func (c *Client) CustomFieldDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointCustomField, id)
}

// CustomFieldUpdate : Update an existing CustomField record in Nautobot.
func (c *Client) CustomFieldUpdate(id uuid.UUID, patch map[string]any) (*types.CustomField, error) {
	return core.Update[types.CustomField](c.Client, extrasEndpointCustomField, id, patch)
}
