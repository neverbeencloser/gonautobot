package tenancy

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/types"
)

const (
	tenancyEndpointTenant = "tenancy/tenants/"
)

type (
	// Tenant : defines a tenant entry in Nautobot
	Tenant struct {
		ID             uuid.UUID      `json:"id"`
		CircuitCount   int            `json:"circuit_count"`
		Comments       string         `json:"comments"`
		Created        time.Time      `json:"created"`
		CustomFields   map[string]any `json:"custom_fields"`
		Description    string         `json:"description"`
		DeviceCount    int            `json:"device_count"`
		Display        string         `json:"display"`
		IpaddressCount int            `json:"ipaddress_count"`
		LastUpdated    time.Time      `json:"last_updated"`
		Name           string         `json:"name"`
		NaturalSlug    string         `json:"natural_slug"`
		NotesURL       string         `json:"notes_url"`
		ObjectType     string         `json:"object_type"`
		PrefixCount    int            `json:"prefix_count"`
		RackCount      int            `json:"rack_count"`
		Tags           []types.Tag    `json:"tags"`
		TenantGroup    *TenantGroup   `json:"tenant_group"`
		URL            string         `json:"url"`
		VMCount        int            `json:"virtualmachine_count"`
		VlanCount      int            `json:"vlan_count"`
		VrfCount       int            `json:"vrf_count"`
	}

	// NewTenant : defines a new tenant entry in Nautobot
	NewTenant struct {
		Name          string         `json:"name"`
		Comments      string         `json:"comments,omitempty"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
		TenantGroup   string         `json:"tenant_group,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
	}
)

// TenantFilter : Go function to process requests for the endpoint: /api/tenancy/tenants/
//
// https://demo.nautobot.com/api/docs/#/tenancy/tenancy_tenants_list
func (c *Client) TenantFilter(q *url.Values) ([]Tenant, error) {
	resp := make([]Tenant, 0)
	return resp, core.Paginate[Tenant](c.Client, tenancyEndpointTenant, q, &resp)
}

// TenantAll : Go function to process requests for the endpoint: /api/tenancy/tenants/
//
// https://demo.nautobot.com/api/docs/#/tenancy/tenancy_tenants_list
func (c *Client) TenantAll() ([]Tenant, error) {
	resp := make([]Tenant, 0)
	return resp, core.Paginate[Tenant](c.Client, tenancyEndpointTenant, nil, &resp)
}

// TenantGet : Get a Tenant by UUID identifier.
func (c *Client) TenantGet(id uuid.UUID) (*Tenant, error) {
	return core.Get[Tenant](c.Client, tenancyEndpointTenant, id)
}

// TenantCreate : Generate a new Tenant record in Nautobot.
func (c *Client) TenantCreate(obj NewTenant) (*Tenant, error) {
	return core.Create[Tenant, NewTenant](c.Client, tenancyEndpointTenant, obj)
}

// TenantDelete : Delete a Tenant by UUID identifier.
func (c *Client) TenantDelete(id uuid.UUID) error {
	return core.Delete(c.Client, tenancyEndpointTenant, id)
}

// TenantUpdate : Update an existing Tenant record in Nautobot.
func (c *Client) TenantUpdate(id uuid.UUID, patch map[string]any) (*Tenant, error) {
	return core.Update[Tenant](c.Client, tenancyEndpointTenant, id, patch)
}
