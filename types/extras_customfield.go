package types

import (
	"time"

	"github.com/google/uuid"
)

type (
	// CustomField : Represents a custom field in Nautobot.
	CustomField struct {
		ID              uuid.UUID  `json:"id"`
		AdvancedAI      bool       `json:"advanced_ai"`
		ContentTypes    []string   `json:"content_types"`
		Created         time.Time  `json:"created"`
		Default         any        `json:"default"`
		Description     string     `json:"description"`
		Display         string     `json:"display"`
		FilterLogic     LabelValue `json:"filter_logic"`
		Grouping        string     `json:"grouping"`
		Key             string     `json:"key"`
		Label           string     `json:"label"`
		LastUpdated     time.Time  `json:"last_updated"`
		NaturalSlug     string     `json:"natural_slug"`
		NotesURL        string     `json:"notes_url"`
		ObjectType      string     `json:"object_type"`
		Required        bool       `json:"required"`
		Type            LabelValue `json:"type"`
		URL             string     `json:"url"`
		ValidationMax   *int       `json:"validation_maximum"`
		ValidationMin   *int       `json:"validation_minimum"`
		ValidationRegex string     `json:"validation_regex"`
		Weight          int        `json:"weight"`
	}

	// NewCustomField : Represents a new custom field to be created in Nautobot.
	NewCustomField struct {
		ContentTypes    []string `json:"content_types"`
		FilterLogic     string   `json:"filter_logic"`
		Key             string   `json:"key"`
		Label           string   `json:"label"`
		Type            string   `json:"type"`
		AdvancedAI      bool     `json:"advanced_ai,omitempty"`
		Default         any      `json:"default,omitempty"`
		Description     string   `json:"description,omitempty"`
		Grouping        string   `json:"grouping,omitempty"`
		Required        bool     `json:"required,omitempty"`
		ValidationMin   *int     `json:"validation_minimum,omitempty"`
		ValidationMax   *int     `json:"validation_maximum,omitempty"`
		ValidationRegex string   `json:"validation_regex,omitempty"`
		Weight          int      `json:"weight,omitempty"`
	}
)
