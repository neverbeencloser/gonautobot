package dcim

import (
	"encoding/json"
	"fmt"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared"
	"github.com/josh-silvas/gonautobot/shared/nested"
	"net/http"
	"net/url"
)

type (
	// Device : Data structure to represent a device record in Nautobot.
	Device struct {
		ID                 string                        `json:"id"`
		AssetTag           string                        `json:"asset_tag"`
		Cluster            *nested.VirtualizationCluster `json:"cluster"`
		Comments           string                        `json:"comments" datastore:",noindex"`
		ConfigContext      map[string]interface{}        `json:"config_context"`
		Created            string                        `json:"created"`
		CustomFields       map[string]interface{}        `json:"custom_fields"`
		DeviceRole         *nested.DeviceRole            `json:"device_role"`
		DeviceType         *nested.DeviceType            `json:"device_type"`
		Display            string                        `json:"display"`
		Face               *shared.LabelValue            `json:"face"`
		LastUpdated        string                        `json:"last_updated"`
		LocalContextSchema *nested.ConfigContextSchema   `json:"local_context_schema"`
		LocalContextData   map[string]interface{}        `json:"local_context_data"`
		Location           *nested.Location              `json:"location"`
		Name               string                        `json:"name"`
		NotesURL           string                        `json:"notes_url"`
		ParentDevice       *Device                       `json:"parent_device"`
		Platform           *nested.Platform              `json:"platform"`
		Position           *int                          `json:"position"`
		PrimaryIP          *nested.IPAddress             `json:"primary_ip"`
		PrimaryIP4         *nested.IPAddress             `json:"primary_ip4"`
		PrimaryIP6         *nested.IPAddress             `json:"primary_ip6"`
		Rack               *nested.Rack                  `json:"rack"`
		SecretsGroup       *nested.SecretsGroup          `json:"secrets_group"`
		Serial             string                        `json:"serial"`
		Site               *Site                         `json:"site"`
		Status             *shared.LabelValue            `json:"status"`
		Tags               []extras.Tag                  `json:"tags"`
		Tenant             *nested.Tenant                `json:"tenant"`
		URL                string                        `json:"url"`
		VCPosition         *int                          `json:"vc_position"`
		VCPriority         *int                          `json:"vc_priority"`
		VirtualChassis     *nested.VirtualChassis        `json:"virtual_chassis"`
	}
)

// GetDevice : Go function to process requests for the endpoint: /api/dcim/devices/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_devices_retrieve
func (c *Client) GetDevice(uuid string, q *url.Values) (*Device, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/devices/%s/", url.PathEscape(uuid)), nil, q)
	if err != nil {
		return nil, err
	}

	ret := new(Device)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// GetDevices : Go function to process requests for the endpoint: /api/dcim/devices/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_devices_list
func (c *Client) GetDevices(query *url.Values) ([]Device, error) {
	devices := make([]Device, 0)
	offset := 0
	if query == nil {
		query = &url.Values{}
	}
	for {
		query.Set("offset", fmt.Sprintf("%d", offset))
		req, err := c.Request(http.MethodGet, "dcim/devices/", nil, query)
		if err != nil {
			return nil, err
		}

		resp := new(shared.ResponseList)
		ret := make([]Device, 0)

		if err = c.UnmarshalDo(req, resp); err != nil {
			return ret, err
		}

		if err = json.Unmarshal(resp.Results, &ret); err != nil {
			return devices, fmt.Errorf("GetDevices.error.json.Unmarshal(%w)", err)
		}
		devices = append(devices, ret...)
		if resp.Count <= len(devices) {
			break
		}
		offset += 50
	}

	return devices, nil
}
