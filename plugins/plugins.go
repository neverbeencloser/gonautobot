package plugins

import (
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/plugins/bgp"
)

// Client : Abstracted base client to use with Nautobot
type Client struct {
	BGP *bgp.Client
}

// New : Initializes the BGP client.
func New(c *core.Client) *Client {
	return &Client{
		BGP: bgp.New(c),
	}
}
