package types

import (
	"time"

	"github.com/google/uuid"
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
		Tags           []Tag          `json:"tags"`
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
