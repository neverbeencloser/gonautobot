package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// LocationType : Represents a location type in Nautobot.
	LocationType struct {
		ID           uuid.UUID      `json:"id"`
		ContentTypes []string       `json:"content_types"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		Nestable     bool           `json:"nestable"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		Parent       *LocationType  `json:"parent"`
		TreeDepth    *int           `json:"tree_depth"`
		URL          string         `json:"url"`
	}

	// NewLocationType : Represents a new location type to be created in Nautobot.
	NewLocationType struct {
		Name         string         `json:"name"`
		ContentTypes []string       `json:"content_types"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
		Description  string         `json:"description,omitempty"`
		Nestable     bool           `json:"nestable,omitempty"`
		Parent       string         `json:"parent,omitempty"`
	}
)
