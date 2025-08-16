package types

import (
	"time"
)

type (
	// Provider : defines a circuit provider in Nautobot
	Provider struct {
		ID           string         `json:"id"`
		Display      string         `json:"display"`
		URL          string         `json:"url"`
		Name         string         `json:"name"`
		Slug         string         `json:"slug"`
		Asn          int            `json:"asn"`
		Account      string         `json:"account"`
		PortalURL    string         `json:"portal_url"`
		NocContact   string         `json:"noc_contact"`
		AdminContact string         `json:"admin_contact"`
		Comments     string         `json:"comments"`
		CircuitCount int            `json:"circuit_count"`
		Created      string         `json:"created"`
		LastUpdated  time.Time      `json:"last_updated"`
		Tags         []Tag          `json:"tags"`
		NotesURL     string         `json:"notes_url"`
		CustomFields map[string]any `json:"custom_fields"`
	}
)
