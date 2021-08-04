package client

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/client/circuits"
	"github.com/josh5276/gonautobot/client/dcim"
	"github.com/josh5276/gonautobot/client/extras"
	"github.com/josh5276/gonautobot/client/graphql"
	"github.com/josh5276/gonautobot/client/ipam"
	"github.com/josh5276/gonautobot/client/plugins"
	"github.com/josh5276/gonautobot/client/status"
	"github.com/josh5276/gonautobot/client/tenancy"
	"github.com/josh5276/gonautobot/client/users"
	"github.com/josh5276/gonautobot/client/virtualization"
)

// Default nautobot API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "demo.nautobot.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new nautobot API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Nautobot {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new nautobot API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Nautobot {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new nautobot API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Nautobot {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Nautobot)
	cli.Transport = transport
	cli.Circuits = circuits.New(transport, formats)
	cli.DCIM = dcim.New(transport, formats)
	cli.Extras = extras.New(transport, formats)
	cli.Graphql = graphql.New(transport, formats)
	cli.IPAM = ipam.New(transport, formats)
	cli.Plugins = plugins.New(transport, formats)
	cli.Status = status.New(transport, formats)
	cli.Tenancy = tenancy.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.Virtualization = virtualization.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Nautobot is a client for nautobot API
type Nautobot struct {
	Circuits       circuits.ClientService
	DCIM           dcim.ClientService
	Extras         extras.ClientService
	Graphql        graphql.ClientService
	IPAM           ipam.ClientService
	Plugins        plugins.ClientService
	Status         status.ClientService
	Tenancy        tenancy.ClientService
	Users          users.ClientService
	Virtualization virtualization.ClientService
	Transport      runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Nautobot) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Circuits.SetTransport(transport)
	c.DCIM.SetTransport(transport)
	c.Extras.SetTransport(transport)
	c.Graphql.SetTransport(transport)
	c.IPAM.SetTransport(transport)
	c.Plugins.SetTransport(transport)
	c.Status.SetTransport(transport)
	c.Tenancy.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.Virtualization.SetTransport(transport)
}
