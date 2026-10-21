package ipam

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/query"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	ipamEndpointPrefix = "ipam/prefixes/"
)

// PrefixGet : Get a Prefix by UUID identifier.
func (c *Client) PrefixGet(id uuid.UUID, opts ...query.Option) (*types.Prefix, error) {
	return core.Get[types.Prefix](c.Client, ipamEndpointPrefix, id, opts...)
}

// PrefixList : Get a list of Prefixes in Nautobot, filtered by any query options.
func (c *Client) PrefixList(opts ...query.Option) ([]types.Prefix, error) {
	prefixes := make([]types.Prefix, 0)
	return prefixes, core.Paginate[types.Prefix](c.Client, ipamEndpointPrefix, opts, &prefixes)
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
func (c *Client) GetPrefixAvailableIPs(id uuid.UUID, opts ...query.Option) ([]types.PrefixAvailableIP, error) {
	q := query.New(opts...)
	req, err := c.Request(
		http.MethodGet,
		ipamEndpointPrefix+fmt.Sprintf("%s/available-ips/", url.PathEscape(id.String())),
		nil,
		&q,
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
func (c *Client) GetPrefixAvailablePrefixes(id uuid.UUID, opts ...query.Option) ([]types.PrefixAvailablePrefix, error) {
	q := query.New(opts...)
	req, err := c.Request(
		http.MethodGet,
		ipamEndpointPrefix+fmt.Sprintf("%s/available-prefixes/", url.PathEscape(id.String())),
		nil,
		&q,
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
