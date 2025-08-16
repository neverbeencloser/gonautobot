package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// SoftwareImageFile represents a software image file in Nautobot.
	SoftwareImageFile struct {
		ID                  uuid.UUID       `json:"id"`
		Created             time.Time       `json:"created"`
		CustomFields        map[string]any  `json:"custom_fields"`
		DefaultImage        bool            `json:"default_image"`
		Display             string          `json:"display"`
		DownloadURL         string          `json:"download_url"`
		ExternalIntegration *Object         `json:"external_integration"`
		HashingAlgorithm    string          `json:"hashing_algorithm"`
		ImageFileChecksum   string          `json:"image_file_checksum"`
		ImageFileName       string          `json:"image_file_name"`
		ImageFileSize       *int            `json:"image_file_size"`
		LastUpdated         time.Time       `json:"last_updated"`
		NaturalSlug         string          `json:"natural_slug"`
		NotesURL            string          `json:"notes_url"`
		ObjectType          string          `json:"object_type"`
		SoftwareVersion     SoftwareVersion `json:"software_version"`
		Status              Status          `json:"status"`
		Tags                []Tag           `json:"tags"`
		URL                 string          `json:"url"`
	}

	// NewSoftwareImageFile represents the data required to create a new software image file.
	NewSoftwareImageFile struct {
		ImageFileName       string         `json:"image_file_name"`
		SoftwareVersion     string         `json:"software_version"`
		Status              string         `json:"status"`
		CustomFields        map[string]any `json:"custom_fields,omitempty"`
		DefaultImage        bool           `json:"default_image"`
		DownloadURL         string         `json:"download_url,omitempty"`
		ExternalIntegration string         `json:"external_integration,omitempty"`
		HashingAlgorithm    string         `json:"hashing_algorithm,omitempty"`
		ImageFileChecksum   string         `json:"image_file_checksum,omitempty"`
		ImageFileSize       int            `json:"image_file_size,omitempty"`
		Relationships       map[string]any `json:"relationships,omitempty"`
		Tags                []string       `json:"tags,omitempty"`
	}
)
