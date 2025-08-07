package extras

import (
	"github.com/neverbeencloser/gonautobot/core"
)

// Client : Abstracted base client to use with Nautobot
type Client struct {
	*core.Client
}

// New : Initializes the extras client.
func New(c *core.Client) *Client {
	return &Client{c}
}
