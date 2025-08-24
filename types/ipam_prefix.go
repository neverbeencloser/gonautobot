package types

import (
	"time"

	"github.com/google/uuid"
)

const (
	// PrefixTypeContainer : Prefix Type Container
	PrefixTypeContainer PrefixType = "container"
	// PrefixTypeNetwork : Prefix Type Network
	PrefixTypeNetwork PrefixType = "network"
	// PrefixTypePool : Prefix Type Pool
	PrefixTypePool PrefixType = "pool"
)

type (
	// Prefix : defines an IPAM Prefix entry in Nautobot
	Prefix struct {
		ID            uuid.UUID      `json:"id"`
		Broadcast     string         `json:"broadcast"`
		Created       time.Time      `json:"created"`
		CustomFields  map[string]any `json:"custom_fields"`
		DateAllocated *time.Time     `json:"date_allocated"`
		Description   string         `json:"description"`
		Display       string         `json:"display"`
		IPVersion     int            `json:"ip_version"`
		LastUpdated   time.Time      `json:"last_updated"`
		Locations     []Location     `json:"location"`
		Namespace     Namespace      `json:"namespace"`
		NaturalSlug   string         `json:"natural_slug"`
		NotesURL      string         `json:"notes_url"`
		ObjectType    string         `json:"object_type"`
		Parent        *Prefix        `json:"parent"`
		Prefix        string         `json:"prefix"`
		PrefixLength  int            `json:"prefix_length"`
		RIR           *Object        `json:"rir"`
		Role          *Role          `json:"role"`
		Status        *Status        `json:"status"`
		Tags          []Tag          `json:"tags"`
		Tenant        *Tenant        `json:"tenant"`
		Type          LabelValue     `json:"type"`
		URL           string         `json:"url"`
		VLAN          *VLAN          `json:"vlan"`
		VRFs          []VRF          `json:"vrf"`
	}

	// NewPrefix : The data structure required to create a new Prefix object in Nautobot.
	NewPrefix struct {
		Prefix        string         `json:"prefix"`
		Status        string         `json:"status"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		DateAllocated string         `json:"date_allocated,omitempty"`
		Description   string         `json:"description,omitempty"`
		Location      string         `json:"location,omitempty"`
		Namespace     string         `json:"namespace,omitempty"`
		Parent        string         `json:"parent,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
		RIR           string         `json:"rir,omitempty"`
		Role          string         `json:"role,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
		Tenant        string         `json:"tenant,omitempty"`
		Type          PrefixType     `json:"type,omitempty"`
		VLAN          string         `json:"vlan,omitempty"`
	}

	// PrefixType : Enum for Prefix Type field
	PrefixType string

	// PrefixAvailableIP : stub IPAddress entry returned by the /prefixes/:id/available-ips/ methods
	PrefixAvailableIP struct {
		Address   string `json:"address"`
		IPVersion int    `json:"ip_version"`
	}

	// PrefixAvailablePrefix : stub Prefix entry returned by the /prefixes/:id/available-prefixes/ methods
	PrefixAvailablePrefix struct {
		IPVersion int    `json:"ip_version"`
		Prefix    string `json:"prefix"`
	}
)
