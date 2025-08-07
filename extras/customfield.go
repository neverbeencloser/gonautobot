package extras

import (
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/core"
	"github.com/neverbeencloser/gonautobot/shared"
)

const (
	extrasEndpointCustomField = "extras/custom-fields/"
)

type (
	// CustomField : Represents a custom field in Nautobot.
	CustomField struct {
		ID              uuid.UUID         `json:"id"`
		AdvancedAI      bool              `json:"advanced_ai"`
		ContentTypes    []string          `json:"content_types"`
		Created         time.Time         `json:"created"`
		Default         any               `json:"default"`
		Description     string            `json:"description"`
		Display         string            `json:"display"`
		FilterLogic     shared.LabelValue `json:"filter_logic"`
		Grouping        string            `json:"grouping"`
		Key             string            `json:"key"`
		Label           string            `json:"label"`
		LastUpdated     time.Time         `json:"last_updated"`
		NaturalSlug     string            `json:"natural_slug"`
		NotesURL        string            `json:"notes_url"`
		ObjectType      string            `json:"object_type"`
		Required        bool              `json:"required"`
		Type            shared.LabelValue `json:"type"`
		URL             string            `json:"url"`
		ValidationMax   *int              `json:"validation_maximum"`
		ValidationMin   *int              `json:"validation_minimum"`
		ValidationRegex string            `json:"validation_regex"`
		Weight          int               `json:"weight"`
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

// CustomFieldGet : Get a CustomField by UUID identifier.
func (c *Client) CustomFieldGet(id uuid.UUID) (*CustomField, error) {
	return core.Get[CustomField](c.Client, extrasEndpointCustomField, id)
}

// CustomFieldFilter : Get a list of CustomFields based on query parameters.
func (c *Client) CustomFieldFilter(q *url.Values) ([]CustomField, error) {
	customFields := make([]CustomField, 0)
	return customFields, core.Paginate[CustomField](c.Client, extrasEndpointCustomField, q, &customFields)
}

// CustomFieldAll : Get all CustomFields in Nautobot.
func (c *Client) CustomFieldAll() ([]CustomField, error) {
	customFields := make([]CustomField, 0)
	return customFields, core.Paginate[CustomField](c.Client, extrasEndpointCustomField, nil, &customFields)
}

// CustomFieldCreate : Generate a new CustomField record in Nautobot.
func (c *Client) CustomFieldCreate(obj NewCustomField) (*CustomField, error) {
	return core.Create[CustomField, NewCustomField](c.Client, extrasEndpointCustomField, obj)
}

// CustomFieldDelete : Delete a CustomField by UUID identifier.
func (c *Client) CustomFieldDelete(id uuid.UUID) error {
	return core.Delete(c.Client, extrasEndpointCustomField, id)
}

// CustomFieldUpdate : Update an existing CustomField record in Nautobot.
func (c *Client) CustomFieldUpdate(id uuid.UUID, patch map[string]any) (*CustomField, error) {
	return core.Update[CustomField](c.Client, extrasEndpointCustomField, id, patch)
}
