package nautobot

import (
	"fmt"
	"github.com/josh-silvas/gonautobot/circuits"
	"github.com/josh-silvas/gonautobot/core"
	"github.com/josh-silvas/gonautobot/dcim"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/ipam"
	"github.com/josh-silvas/gonautobot/plugins"
	"github.com/josh-silvas/gonautobot/tenancy"
	"github.com/josh-silvas/gonautobot/virtualization"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type (
	// Client : Stored memory objects for the Nautobot client.
	Client struct {
		Request *core.Client

		Circuits       *circuits.Client
		Dcim           *dcim.Client
		Extras         *extras.Client
		Ipam           *ipam.Client
		Plugins        *plugins.Client
		Tenancy        *tenancy.Client
		Virtualization *virtualization.Client
	}

	// Option : Basic options allowed with this client.
	Option func(*Client)
)

// WithEndpoint : Sets the API endpoint to use in Nautobot. Default is https://demo.nautobot.com
func WithEndpoint(url string) Option {
	return func(o *Client) {
		o.Request.URL = SanitizeURL(url)
	}
}

// WithHTTPClient : Overrides the default HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(o *Client) {
		o.Request.Client = client
	}
}

// WithLogger : Overrides the default logger for the package.
func WithLogger(logger *logrus.Logger) Option {
	return func(o *Client) {
		o.Request.Log = logger
	}
}

// New : Function used to create a new Nautobot client data type.
func New(token string, opts ...Option) *Client {
	c := &Client{
		Request: &core.Client{
			Token:  token,
			URL:    "https://demo.nautobot.com/api/",
			Client: http.DefaultClient,
			Log:    logrus.StandardLogger(),
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Circuits = circuits.New(c.Request)
	c.Dcim = dcim.New(c.Request)
	c.Extras = extras.New(c.Request)
	c.Ipam = ipam.New(c.Request)
	c.Plugins = plugins.New(c.Request)
	c.Tenancy = tenancy.New(c.Request)
	c.Virtualization = virtualization.New(c.Request)
	return c
}

// SanitizeURL : Sanitizes a URL for use with Nautobot.
func SanitizeURL(s string) string {
	if !strings.HasPrefix(s, "http://") && !strings.HasPrefix(s, "https://") {
		s = fmt.Sprintf("https://%s", s)
	}
	if !strings.HasSuffix(s, "/") {
		s = fmt.Sprintf("%s/", s)
	}
	if !strings.HasSuffix(s, "api/") {
		s = fmt.Sprintf("%sapi/", s)
	}
	return s
}
