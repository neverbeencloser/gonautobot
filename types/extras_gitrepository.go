package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/types/nested"
)

type (
	// GitRepository : Data type entry for a Git Repository in Nautobot.
	GitRepository struct {
		ID               uuid.UUID            `json:"id"`
		ObjectType       string               `json:"object_type"`
		Display          string               `json:"display"`
		URL              string               `json:"url"`
		NaturalSlug      string               `json:"natural_slug"`
		ProvidedContents []string             `json:"provided_contents"`
		Name             string               `json:"name"`
		Slug             string               `json:"slug"`
		RemoteURL        string               `json:"remote_url"`
		Branch           string               `json:"branch"`
		CurrentHead      string               `json:"current_head"`
		SecretsGroup     *nested.SecretsGroup `json:"secrets_group"`
		Created          time.Time            `json:"created"`
		LastUpdated      time.Time            `json:"last_updated"`
		NotesURL         string               `json:"notes_url"`
		CustomFields     map[string]any       `json:"custom_fields"`
		ComputedFields   map[string]any       `json:"computed_fields"`
		Relationships    map[string]any       `json:"relationships"`
		Tags             []Tag                `json:"tags"`
	}

	// NewGitRepository : Structured input for a new Git Repository record in Nautobot.
	NewGitRepository struct {
		Name             string         `json:"name"`
		RemoteURL        string         `json:"remote_url"`
		Slug             string         `json:"slug,omitempty"`
		Branch           string         `json:"branch,omitempty"`
		ProvidedContents []string       `json:"provided_contents,omitempty"`
		CurrentHead      string         `json:"current_head,omitempty"`
		SecretsGroup     string         `json:"secrets_group,omitempty"`
		CustomFields     map[string]any `json:"custom_fields,omitempty"`
		Relationships    map[string]any `json:"relationships,omitempty"`
		Tags             []string       `json:"tags,omitempty"`
	}
)
