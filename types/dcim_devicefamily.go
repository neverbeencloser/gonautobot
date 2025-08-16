package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// DeviceFamily represents a device family in Nautobot's DCIM module.
	DeviceFamily struct {
		ID           uuid.UUID      `json:"id"`
		Created      time.Time      `json:"created"`
		CustomFields map[string]any `json:"custom_fields"`
		Description  string         `json:"description"`
		Display      string         `json:"display"`
		LastUpdated  time.Time      `json:"last_updated"`
		Name         string         `json:"name"`
		NaturalSlug  string         `json:"natural_slug"`
		NotesURL     string         `json:"notes_url"`
		ObjectType   string         `json:"object_type"`
		Tags         []Tag          `json:"tags"`
		URL          string         `json:"url"`
	}

	// NewDeviceFamily represents the data required to create a new device family.
	NewDeviceFamily struct {
		Name         string         `json:"name"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
		Description  string         `json:"description,omitempty"`
		Tags         []string       `json:"tags,omitempty"`
	}
)
