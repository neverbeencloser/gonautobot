package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// VLAN : defines a vlan entry in Nautobot
	VLAN struct {
		ID           uuid.UUID      `json:"id"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Locations    []Location     `json:"location"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		PrefixCount  int            `json:"prefix_count"`
		Role         *Role          `json:"role"`
		Status       Status         `json:"status"`
		Tags         []Tag          `json:"tags"`
		Tenant       *Tenant        `json:"tenant"`
		URL          string         `json:"url"`
		VID          int            `json:"vid"`
		VLANGroup    *Object        `json:"vlan_group"`
	}

	// NewVLAN : defines the information needed to create a new VLAN in Nautobot
	NewVLAN struct {
		Name          string         `json:"name"`
		Status        string         `json:"status"`
		VID           int            `json:"vid"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
		Role          string         `json:"role,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
		Tenant        string         `json:"tenant,omitempty"`
		VLANGroup     string         `json:"vlan_group,omitempty"`
	}
)
