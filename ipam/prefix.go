package ipam

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	ipamEndpointPrefix = "ipam/prefixes/"
)

// PrefixGet : Get a Prefix by UUID identifier.
func (c *Client) PrefixGet(id uuid.UUID) (*types.Prefix, error) {
	return core.Get[types.Prefix](c.Client, ipamEndpointPrefix, id)
}

// PrefixFilter : Get a list of Prefixes based on query parameters.
func (c *Client) PrefixFilter(q *url.Values) ([]types.Prefix, error) {
	prefixes := make([]types.Prefix, 0)
	return prefixes, core.Paginate[types.Prefix](c.Client, ipamEndpointPrefix, q, &prefixes)
}

// PrefixAll : Get all Prefixes in Nautobot.
func (c *Client) PrefixAll() ([]types.Prefix, error) {
	prefixes := make([]types.Prefix, 0)
	return prefixes, core.Paginate[types.Prefix](c.Client, ipamEndpointPrefix, nil, &prefixes)
}

// PrefixCreate : Generate a new Prefix record in Nautobot.
func (c *Client) PrefixCreate(obj types.NewPrefix) (*types.Prefix, error) {
	return core.Create[types.Prefix, types.NewPrefix](c.Client, ipamEndpointPrefix, obj)
}

// PrefixUpdate : Update an existing Prefix record in Nautobot.
func (c *Client) PrefixUpdate(id uuid.UUID, patch map[string]any) (*types.Prefix, error) {
	return core.Update[types.Prefix](c.Client, ipamEndpointPrefix, id, patch)
}

// PrefixDelete : Delete a Prefix by UUID identifier.
func (c *Client) PrefixDelete(id uuid.UUID) error {
	return core.Delete(c.Client, ipamEndpointPrefix, id)
}

// GetPrefixAvailableIPs : Go function to process requests for the endpoint: /api/ipam/prefixes/:id/available-ips/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_available_ips_list
func (c *Client) GetPrefixAvailableIPs(id uuid.UUID, q *url.Values) ([]types.PrefixAvailableIP, error) {
	req, err := c.Request(
		http.MethodGet,
		ipamEndpointPrefix+fmt.Sprintf("%s/available-ips/", url.PathEscape(id.String())),
		nil,
		q,
	)
	if err != nil {
		return nil, err
	}

	ret := make([]types.PrefixAvailableIP, 0)
	if err := c.UnmarshalDo(req, &ret); err != nil {
		return nil, fmt.Errorf("GetPrefixAvailableIPs.error.UnmarshalDo(%w)", err)
	}
	return ret, nil
}

// GetPrefixAvailablePrefixes : Go function to process requests for the endpoint: /api/ipam/prefixes/:id/available-prefixes/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_available_prefixes_list
func (c *Client) GetPrefixAvailablePrefixes(id uuid.UUID, q *url.Values) ([]types.PrefixAvailablePrefix, error) {
	req, err := c.Request(
		http.MethodGet,
		ipamEndpointPrefix+fmt.Sprintf("%s/available-prefixes/", url.PathEscape(id.String())),
		nil,
		q,
	)
	if err != nil {
		return nil, err
	}

	ret := make([]types.PrefixAvailablePrefix, 0)
	if err := c.UnmarshalDo(req, &ret); err != nil {
		return nil, fmt.Errorf("GetPrefixAvailablePrefixes.error.UnmarshalDo(%w)", err)
	}
	return ret, nil
}
