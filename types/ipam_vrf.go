package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// VRF : Data type entry for a VRF in Nautobot.
	VRF struct {
		ID                    uuid.UUID      `json:"id"`
		Created               time.Time      `json:"created"`
		CustomFields          map[string]any `json:"custom_fields"`
		Description           string         `json:"description"`
		Devices               []Device       `json:"devices"`
		Display               string         `json:"display"`
		ExportTargets         []Object       `json:"export_targets"`
		ImportTargets         []Object       `json:"import_targets"`
		LastUpdated           time.Time      `json:"last_updated"`
		Name                  string         `json:"name"`
		Namespace             Namespace      `json:"namespace"`
		NaturalSlug           string         `json:"natural_slug"`
		NotesURL              string         `json:"notes_url"`
		ObjectType            string         `json:"object_type"`
		Prefixes              []Prefix       `json:"prefixes"`
		RD                    *string        `json:"rd"`
		Status                *Status        `json:"status"`
		Tags                  []Tag          `json:"tags"`
		Tenant                *Tenant        `json:"tenant"`
		URL                   string         `json:"url"`
		VirtualDeviceContexts []Object       `json:"virtual_device_contexts"`
		VirtualMachines       []Object       `json:"virtual_machines"`
	}

	// NewVRF : Structured input for a new VRF record in Nautobot.
	NewVRF struct {
		Name          string         `json:"name"`
		Namespace     string         `json:"namespace"`
		RD            string         `json:"rd"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		Description   string         `json:"description,omitempty"`
		ExportTargets []string       `json:"export_targets,omitempty"`
		ImportTargets []string       `json:"import_targets,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
		Status        string         `json:"status,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
		Tenant        string         `json:"tenant,omitempty"`
	}
)
