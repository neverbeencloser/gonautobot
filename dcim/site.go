package dcim

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/neverbeencloser/gonautobot/types"
)

// SiteGet : Go function to process requests for the endpoint: /api/dcim/sites/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_sites_retrieve
func (c *Client) SiteGet(uuid string) (*types.Site, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/sites/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(types.Site)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// SiteFilter : Go function to process requests for the endpoint: /api/dcim/sites/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_sites_list
func (c *Client) SiteFilter(q *url.Values) ([]types.Site, error) {
	req, err := c.Request(http.MethodGet, "dcim/sites/", nil, q)
	if err != nil {
		return nil, err
	}

	resp := new(types.ResponseList)
	ret := make([]types.Site, 0)

	if err = c.UnmarshalDo(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("GetSites.error.json.Unmarshal(%w)", err)
	}
	return ret, err
}
