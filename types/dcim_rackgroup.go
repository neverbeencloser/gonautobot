package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// RackGroup : Data type entry for a RackGroup in Nautobot.
	RackGroup struct {
		ID           uuid.UUID      `json:"id"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Location     Location       `json:"location"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		Parent       *RackGroup     `json:"parent"`
		RackCount    int            `json:"rack_count"`
		TreeDepth    *int           `json:"tree_depth"`
		URL          string         `json:"url"`
	}

	// NewRackGroup represents the structure for creating a new RackGroup in Nautobot.
	NewRackGroup struct {
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		Location      string         `json:"location"`
		Name          string         `json:"name"`
		Parent        string         `json:"parent,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
	}
)
