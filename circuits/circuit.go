package circuits

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

// CircuitGet : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_list
func (c *Client) CircuitGet(uuid string) (types.Circuit, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("circuits/circuits/%s/", uuid), nil, nil)
	if err != nil {
		return types.Circuit{}, err
	}

	ret := new(types.Circuit)
	return *ret, c.UnmarshalDo(req, ret)
}

// CircuitFilter : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_list
func (c *Client) CircuitFilter(q *url.Values) ([]types.Circuit, error) {
	resp := make([]types.Circuit, 0)
	return resp, core.Paginate[types.Circuit](c.Client, "circuits/circuits/", q, &resp)
}

// CircuitAll : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_list
func (c *Client) CircuitAll() ([]types.Circuit, error) {
	resp := make([]types.Circuit, 0)
	return resp, core.Paginate[types.Circuit](c.Client, "circuits/circuits/", nil, &resp)
}

// CircuitDelete : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_destroy
func (c *Client) CircuitDelete(uuid string) error {
	req, err := c.Request(http.MethodDelete, fmt.Sprintf("circuits/circuits/%s/", uuid), nil, nil)
	if err != nil {
		return err
	}
	return c.UnmarshalDo(req, nil)
}

// CircuitCreate : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_create
func (c *Client) CircuitCreate(r types.CircuitRequest) (types.Circuit, error) {
	req, err := c.Request(http.MethodPost, "circuits/circuits/", r, nil)
	if err != nil {
		return types.Circuit{}, err
	}
	var ret types.Circuit
	return ret, c.UnmarshalDo(req, &ret)
}
