package tenancy

import (
	"encoding/json"
	"fmt"
	"github.com/josh-silvas/gonautobot/shared"
	"net/http"
	"net/url"
	"time"
)

type (
	// TenantGroup : defines a tenant-group entry in Nautobot
	TenantGroup struct {
		ID           string         `json:"id"`
		Display      string         `json:"display"`
		URL          string         `json:"url"`
		Name         string         `json:"name"`
		Slug         string         `json:"slug"`
		Parent       *TenantGroup   `json:"parent"`
		Description  string         `json:"description"`
		TenantCount  int            `json:"tenant_count"`
		Depth        int            `json:"_depth"`
		Created      string         `json:"created"`
		LastUpdated  time.Time      `json:"last_updated"`
		NotesURL     string         `json:"notes_url"`
		CustomFields map[string]any `json:"custom_fields"`
	}
)

// GetTenantGroups : Go function to process requests for the endpoint: /api/tenancy/tenant-groups/
//
// https://demo.nautobot.com/api/docs/#/tenancy/tenancy_tenant_groups_list
func (c *Client) GetTenantGroups(q *url.Values) ([]TenantGroup, error) {
	req, err := c.Request(http.MethodGet, "tenancy/tenant-groups/", nil, q)
	if err != nil {
		return nil, err
	}

	resp := new(shared.ResponseList)
	ret := make([]TenantGroup, 0)
	if err = c.UnmarshalDo(req, resp); err != nil {
		return ret, err
	}

	if err = json.Unmarshal(resp.Results, &ret); err != nil {
		err = fmt.Errorf("GetTenantGroups.error.json.Unmarshal(%w)", err)
	}
	return ret, err
}
