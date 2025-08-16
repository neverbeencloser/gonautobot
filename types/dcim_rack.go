package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// Rack : Data type entry for a Rack in Nautobot.
	Rack struct {
		ID             uuid.UUID      `json:"id"`
		AssetTag       *string        `json:"asset_tag"`
		Comments       string         `json:"comments"`
		Created        time.Time      `json:"created"`
		CustomFields   map[string]any `json:"custom_fields"`
		DescUnits      bool           `json:"desc_units"`
		DeviceCount    int            `json:"device_count"`
		Display        string         `json:"display"`
		FacilityID     *string        `json:"facility_id"`
		LastUpdated    time.Time      `json:"last_updated"`
		Location       Location       `json:"location"`
		Name           string         `json:"name"`
		NaturalSlug    string         `json:"natural_slug"`
		NotesURL       string         `json:"notes_url"`
		ObjectType     string         `json:"object_type"`
		OuterDepth     *int           `json:"outer_depth"`
		OuterUnit      *LabelValue    `json:"outer_unit"`
		OuterWidth     *int           `json:"outer_width"`
		PowerFeedCount int            `json:"power_feed_count"`
		RackGroup      *RackGroup     `json:"rack_group"`
		Role           *Role          `json:"role"`
		Serial         string         `json:"serial"`
		Status         Status         `json:"status"`
		Tags           []Tag          `json:"tags"`
		Tenant         *Tenant        `json:"tenant"`
		Type           *LabelValue    `json:"type"`
		UHeight        int            `json:"u_height"`
		URL            string         `json:"url"`
		Width          LabelValueInt  `json:"width"`
	}

	// NewRack represents the structure for creating a new Rack in Nautobot.
	NewRack struct {
		Location      string         `json:"location"`
		Name          string         `json:"name"`
		Status        string         `json:"status"`
		Width         int            `json:"width"`
		UHeight       int            `json:"u_height"`
		AssetTag      string         `json:"asset_tag,omitempty"`
		Comments      string         `json:"comments,omitempty"`
		CustomFields  map[string]any `json:"custom_fields,omitempty"`
		DescUnits     bool           `json:"desc_units,omitempty"`
		FacilityID    string         `json:"facility_id,omitempty"`
		OuterDepth    int            `json:"outer_depth,omitempty"`
		OuterUnit     string         `json:"outer_unit,omitempty"`
		OuterWidth    int            `json:"outer_width,omitempty"`
		RackGroup     string         `json:"rack_group,omitempty"`
		Relationships map[string]any `json:"relationships,omitempty"`
		Role          string         `json:"role,omitempty"`
		Serial        string         `json:"serial,omitempty"`
		Tags          []string       `json:"tags,omitempty"`
		Tenant        string         `json:"tenant,omitempty"`
		Type          string         `json:"type,omitempty"`
	}
)
