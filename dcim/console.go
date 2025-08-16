package dcim

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/neverbeencloser/gonautobot/types"
)

// ConsoleConnectionFilter : Method used to fetch a list of console connections for a device.
func (c *Client) ConsoleConnectionFilter(q *url.Values) ([]types.ConsolePort, error) {
	req, err := c.Request(http.MethodGet, "dcim/console-connections/", nil, q)
	if err != nil {
		return nil, err
	}

	resp := new(types.ResponseList)
	ret := make([]types.ConsolePort, 0)
	if err = c.UnmarshalDo(req, resp); err != nil {
		return ret, err
	}

	if e := json.Unmarshal(resp.Results, &ret); e != nil { // nolint: musttag
		return nil, fmt.Errorf("ConsoleConnections.error.json.Unmarshal(%w)", e)
	}

	for k, v := range ret {
		ce := strings.ReplaceAll(v.ConnectedEndpoint.Name, " ", "")
		if p, err := strconv.Atoi(strings.Replace(ce, "c", "", -1)); err == nil {
			ret[k].ConnectedEndpoint.Port = p
		}
	}

	return ret, nil
}
