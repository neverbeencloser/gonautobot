package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/types/nested"
)

type (
	// JobHook : A job hook triggers a job receiver on create/update/delete of matching content types.
	JobHook struct {
		ID             uuid.UUID      `json:"id"`
		ObjectType     string         `json:"object_type"`
		Display        string         `json:"display"`
		URL            string         `json:"url"`
		NaturalSlug    string         `json:"natural_slug"`
		ContentTypes   []string       `json:"content_types"`
		Enabled        bool           `json:"enabled"`
		Name           string         `json:"name"`
		TypeCreate     bool           `json:"type_create"`
		TypeDelete     bool           `json:"type_delete"`
		TypeUpdate     bool           `json:"type_update"`
		Job            *nested.Job    `json:"job"`
		Created        time.Time      `json:"created"`
		LastUpdated    time.Time      `json:"last_updated"`
		NotesURL       string         `json:"notes_url"`
		CustomFields   map[string]any `json:"custom_fields"`
		ComputedFields map[string]any `json:"computed_fields"`
		Relationships  map[string]any `json:"relationships"`
	}

	// NewJobHook : Input for creating a JobHook (POST).
	NewJobHook struct {
		Name          string         `json:"name"`
		ContentTypes  []string       `json:"content_types"`
		Job           string         `json:"job"`
		Enabled       *bool          `json:"enabled,omitempty"`
		TypeCreate    *bool          `json:"type_create,omitempty"`
		TypeDelete    *bool          `json:"type_delete,omitempty"`
		TypeUpdate    *bool          `json:"type_update,omitempty"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
	}
)
