package dcim

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

// DeviceGet : Go function to process requests for the endpoint: /api/dcim/devices/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_devices_retrieve
func (c *Client) DeviceGet(uuid string) (*types.Device, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/devices/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(types.Device)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// DeviceFilter : Go function to process requests for the endpoint: /api/dcim/devices/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_devices_list
func (c *Client) DeviceFilter(q *url.Values) ([]types.Device, error) {
	devices := make([]types.Device, 0)
	return devices, core.Paginate[types.Device](c.Client, "dcim/devices/", q, &devices)
}

// DeviceAll : Go function to process requests for the endpoint: /api/dcim/devices/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_devices_list
func (c *Client) DeviceAll() ([]types.Device, error) {
	devices := make([]types.Device, 0)
	return devices, core.Paginate[types.Device](c.Client, "dcim/devices/", nil, &devices)
}

// DeviceCreate : Generate a new Device record in Nautobot.
func (c *Client) DeviceCreate(obj types.NewDevice) (types.Device, error) {
	var r types.Device
	req, err := c.Request(http.MethodPost, "dcim/devices/", obj, nil)
	if err != nil {
		return r, err
	}

	if err := c.UnmarshalDo(req, &r); err != nil {
		return r, fmt.Errorf("DeviceCreate.error.UnmarshalDo(%w)", err)
	}
	return r, nil
}

// DeviceDelete : DCIM method to delete a Device by UUIDv4 identifier.
func (c *Client) DeviceDelete(uuid string) error {
	if uuid == "" {
		return errors.New("DeviceDelete.error.UUID(UUIDv4 is missing)")
	}
	req, err := c.Request(http.MethodDelete, fmt.Sprintf("dcim/devices/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return err
	}
	return c.UnmarshalDo(req, nil)
}

// DeviceUpdate : Update an existing Device record in Nautobot.
func (c *Client) DeviceUpdate(uuid string, patch map[string]any) (types.Device, error) {
	var r types.Device
	req, err := c.Request(http.MethodPatch, fmt.Sprintf("dcim/devices/%s/", uuid), patch, nil)
	if err != nil {
		return r, err
	}
	if err := c.UnmarshalDo(req, &r); err != nil {
		return r, fmt.Errorf("DeviceUpdate.error.UnmarshalDo(%w)", err)
	}
	return r, nil
}
