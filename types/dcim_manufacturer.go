package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Manufacturer represents a manufacturer in Nautobot's DCIM application.
	Manufacturer struct {
		ID                 uuid.UUID      `json:"id"`
		CloudAccountCount  int            `json:"cloud_account_count"`
		Created            time.Time      `json:"created"`
		CustomFields       map[string]any `json:"custom_fields"`
		Description        string         `json:"description"`
		DeviceTypeCount    int            `json:"device_type_count"`
		Display            string         `json:"display"`
		InventoryItemCount int            `json:"inventory_item_count"`
		LastUpdated        time.Time      `json:"last_updated"`
		Name               string         `json:"name"`
		NaturalSlug        string         `json:"natural_slug"`
		NotesURL           string         `json:"notes_url"`
		ObjectType         string         `json:"object_type"`
		PlatformCount      int            `json:"platform_count"`
		URL                string         `json:"url"`
	}

	// NewManufacturer represents the data required to create a new manufacturer.
	NewManufacturer struct {
		Name         string         `json:"name"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
		Description  string         `json:"description,omitempty"`
	}
)
