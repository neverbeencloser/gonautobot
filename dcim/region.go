package dcim

import (
	"encoding/json"
	"fmt"
	"github.com/josh-silvas/gonautobot/shared"
	"net/http"
	"net/url"
)

type (
	// Region : defines a region object in Nautobot
	Region struct {
		ID           string         `json:"id"`
		Created      string         `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Depth        int            `json:"_depth"`
		Description  string         `json:"description" datastore:",noindex"`
		Display      string         `json:"display"`
		LastUpdated  string         `json:"last_updated"`
		Name         string         `json:"name"`
		NotesURL     string         `json:"notes_url"`
		Parent       *Region        `json:"parent"`
		SiteCount    int            `json:"site_count"`
		Slug         string         `json:"slug"`
		URL          string         `json:"url"`
	}
)

// GetRegion : Go function to process requests for the endpoint: /api/dcim/regions/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_regions_retrieve
func (c *Client) GetRegion(uuid string, q *url.Values) (*Region, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/regions/%s/", url.PathEscape(uuid)), nil, q)
	if err != nil {
		return nil, err
	}

	ret := new(Region)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// GetRegions : Go function to process requests for the endpoint: /api/dcim/regions/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_regions_list
func (c *Client) GetRegions(q *url.Values) ([]Region, error) {
	req, err := c.Request(http.MethodGet, "dcim/regions/", nil, q)
	if err != nil {
		return nil, err
	}

	resp := new(shared.ResponseList)
	ret := make([]Region, 0)

	if err = c.UnmarshalDo(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("GetRegions.error.json.Unmarhsal(%w)", err)
	}
	return ret, err
}
