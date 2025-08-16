package dcim

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

// RegionGet : Go function to process requests for the endpoint: /api/dcim/regions/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_regions_retrieve
func (c *Client) RegionGet(uuid string) (*types.Region, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/regions/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(types.Region)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// RegionFilter : Go function to process requests for the endpoint: /api/dcim/regions/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_regions_list
func (c *Client) RegionFilter(q *url.Values) ([]types.Region, error) {
	resp := make([]types.Region, 0)
	if err := core.Paginate[types.Region](c.Client, "dcim/regions/", q, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
