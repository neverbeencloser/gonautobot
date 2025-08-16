package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/types/nested"
)

type (
	// Namespace : Represents a namespace in Nautobot.
	Namespace struct {
		ID           uuid.UUID        `json:"id"`
		Created      time.Time        `json:"created"`
		CustomFields map[string]any   `json:"custom_fields"`
		Description  string           `json:"description"`
		Display      string           `json:"display"`
		LastUpdated  time.Time        `json:"last_updated"`
		Location     *nested.Location `json:"location"`
		Name         string           `json:"name"`
		NaturalSlug  string           `json:"natural_slug"`
		NotesURL     string           `json:"notes_url"`
		ObjectType   string           `json:"object_type"`
		Tags         []Tag            `json:"tags"`
		URL          string           `json:"url"`
	}

	// NewNamespace : Represents a new namespace to be created in Nautobot.
	NewNamespace struct {
		Name          string         `json:"name"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		Location      string         `json:"location,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
	}
)
