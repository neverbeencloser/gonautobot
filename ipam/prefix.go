package ipam

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

// PrefixGet : Go function to process requests for the endpoint: /api/ipam/prefixes/:id/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_retrieve
func (c *Client) PrefixGet(uuid string) (*types.Prefix, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("ipam/prefixes/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(types.Prefix)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// PrefixFilter : Go function to process requests for the endpoint: /api/ipam/prefixes/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_list
func (c *Client) PrefixFilter(q *url.Values) ([]types.Prefix, error) {
	resp := make([]types.Prefix, 0)
	return resp, core.Paginate[types.Prefix](c.Client, "ipam/prefixes/", q, &resp)
}

// PrefixCreate : Creates a new prefix using the NewPrefix data type.
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_create
func (c *Client) PrefixCreate(prefix *types.NewPrefix) (*types.Prefix, error) {
	req, err := c.Request(http.MethodPost, "ipam/prefixes/", prefix, nil)
	if err != nil {
		return nil, err
	}

	var r types.Prefix
	err = c.UnmarshalDo(req, &r)
	if err != nil {
		return nil, fmt.Errorf("CreatePrefix.error.UnmarshalDo(%w)", err)
	}

	return &r, nil
}

// PrefixUpdate : Updates a Nautobot prefix by UUID
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_partial_update
func (c *Client) PrefixUpdate(uuid string, prefix *types.PrefixUpdate) (*types.Prefix, error) {
	req, err := c.Request(http.MethodPatch, fmt.Sprintf("ipam/prefixes/%s/", uuid), prefix, nil)
	if err != nil {
		return nil, err
	}

	var r types.Prefix
	if err := c.UnmarshalDo(req, &r); err != nil {
		return nil, fmt.Errorf("UpdatePrefix.error.UnmarshalDo(%w)", err)
	}

	return &r, nil
}

// PrefixDelete : Removes a prefix from the Nautobot DB.
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_destroy
func (c *Client) PrefixDelete(prefixID string) error {
	req, err := c.Request(http.MethodDelete, fmt.Sprintf("ipam/prefixes/%s/", prefixID), nil, nil)
	if err != nil {
		return err
	}

	if err := c.UnmarshalDo(req, nil); err != nil {
		return fmt.Errorf("DeletePrefix.error.UnmarshalDo(%w)", err)
	}
	return nil
}

// GetPrefixAvailableIPs : Go function to process requests for the endpoint: /api/ipam/prefixes/:id/available-ips/
//
// https://demo.nautobot.com/api/docs/#/ipam/ipam_prefixes_available_ips_list
func (c *Client) GetPrefixAvailableIPs(uuid string, q *url.Values) ([]types.PrefixAvailableIP, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("ipam/prefixes/%s/available-ips/", url.PathEscape(uuid)), nil, q)
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
func (c *Client) GetPrefixAvailablePrefixes(uuid string, q *url.Values) ([]types.PrefixAvailablePrefix, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("ipam/prefixes/%s/available-prefixes/", url.PathEscape(uuid)), nil, q)
	if err != nil {
		return nil, err
	}

	ret := make([]types.PrefixAvailablePrefix, 0)
	if err := c.UnmarshalDo(req, &ret); err != nil {
		return nil, fmt.Errorf("GetPrefixAvailablePrefixes.error.UnmarshalDo(%w)", err)
	}
	return ret, nil
}
