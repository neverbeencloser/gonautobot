package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Role : Data type entry for a Role in Nautobot.
	Role struct {
		ID           uuid.UUID      `json:"id"`
		Color        string         `json:"color"`
		ContentTypes []string       `json:"content_types"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		URL          string         `json:"url"`
		Weight       *int           `json:"weight"`
	}

	// NewRole : Structured input for a new Role record in Nautobot.
	NewRole struct {
		Name         string         `json:"name"`
		ContentTypes []string       `json:"content_types"`
		Color        string         `json:"color,omitempty"`
		Description  string         `json:"description,omitempty"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
		Weight       int            `json:"weight,omitempty"`
	}
)
