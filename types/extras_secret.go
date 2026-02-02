package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Secret : Data type entry for a Secret in Nautobot.
	Secret struct {
		ID             uuid.UUID      `json:"id"`
		ObjectType     string         `json:"object_type"`
		Display        string         `json:"display"`
		URL            string         `json:"url"`
		NaturalSlug    string         `json:"natural_slug"`
		Name           string         `json:"name"`
		Description    string         `json:"description"`
		Provider       string         `json:"provider"`
		Parameters     map[string]any `json:"parameters"`
		Created        time.Time      `json:"created"`
		LastUpdated    time.Time      `json:"last_updated"`
		Tags           []Tag          `json:"tags"`
		NotesURL       string         `json:"notes_url"`
		CustomFields   map[string]any `json:"custom_fields"`
		ComputedFields map[string]any `json:"computed_fields"`
		Relationships  map[string]any `json:"relationships"`
	}

	// NewSecret : Structured input for a new Secret record in Nautobot.
	NewSecret struct {
		Name          string         `json:"name"`
		Provider      string         `json:"provider"`
		Description   string         `json:"description,omitempty"`
		Parameters    map[string]any `json:"parameters,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
	}
)
