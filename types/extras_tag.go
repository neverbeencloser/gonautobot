package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Tag : Data type entry for a Tag in Nautobot.
	Tag struct {
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
		TaggedItems  int            `json:"tagged_items"`
		URL          string         `json:"url"`
	}

	// NewTag : Data type for creating a new Tag in Nautobot.
	NewTag struct {
		Name          string         `json:"name"`
		ContentTypes  []string       `json:"content_types"`
		Color         string         `json:"color,omitempty"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
	}
)
