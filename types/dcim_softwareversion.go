package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// SoftwareVersion represents a software version in Nautobot.
	SoftwareVersion struct {
		ID               uuid.UUID      `json:"id"`
		Alias            string         `json:"alias"`
		Created          time.Time      `json:"created"`
		CustomFields     map[string]any `json:"custom_fields"`
		Display          string         `json:"display"`
		DocumentationURL string         `json:"documentation_url"`
		EndOfSupportDate *string        `json:"end_of_support_date"`
		LastUpdated      time.Time      `json:"last_updated"`
		LongTermSupport  bool           `json:"long_term_support"`
		NaturalSlug      string         `json:"natural_slug"`
		NotesURL         string         `json:"notes_url"`
		ObjectType       string         `json:"object_type"`
		Platform         Platform       `json:"platform"`
		PreRelease       bool           `json:"pre_release"`
		ReleaseDate      *string        `json:"release_date"`
		Status           Status         `json:"status"`
		Tags             []Tag          `json:"tags"`
		URL              string         `json:"url"`
		Version          string         `json:"version"`
	}

	// NewSoftwareVersion represents the data required to create a new software version.
	NewSoftwareVersion struct {
		Platform         string         `json:"platform"`
		Status           string         `json:"status"`
		Version          string         `json:"version"`
		Alias            string         `json:"alias,omitempty"`
		CustomFields     map[string]any `json:"custom_fields,omitempty"`
		DocumentationURL string         `json:"documentation_url,omitempty"`
		EndOfSupportDate string         `json:"end_of_support_date,omitempty"`
		LongTermSupport  bool           `json:"long_term_support"`
		PreRelease       bool           `json:"pre_release"`
		ReleaseDate      string         `json:"release_date,omitempty"`
		Tags             []string       `json:"tags,omitempty"`
	}
)
