package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Relationship relationship
//
// swagger:model Relationship
type Relationship struct {

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Destination filter
	//
	// Queryset filter matching the applicable destination objects of the selected type
	DestinationFilter *string `json:"destination_filter,omitempty"`

	// Hide for destination object
	//
	// Hide this relationship on the destination object.
	DestinationHidden bool `json:"destination_hidden,omitempty"`

	// Destination Label
	//
	// Name of the relationship as displayed on the destination object.
	// Max Length: 50
	DestinationLabel string `json:"destination_label,omitempty"`

	// Destination type
	// Required: true
	DestinationType *string `json:"destination_type"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name
	//
	// Internal relationship name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Slug
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// Source filter
	//
	// Queryset filter matching the applicable source objects of the selected type
	SourceFilter *string `json:"source_filter,omitempty"`

	// Hide for source object
	//
	// Hide this relationship on the source object.
	SourceHidden bool `json:"source_hidden,omitempty"`

	// Source Label
	//
	// Name of the relationship as displayed on the source object.
	// Max Length: 50
	SourceLabel string `json:"source_label,omitempty"`

	// Source type
	// Required: true
	SourceType *string `json:"source_type"`

	// Type
	// Enum: [one-to-one one-to-many many-to-many]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this relationship
func (m *Relationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Relationship) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateDestinationLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationLabel) { // not required
		return nil
	}

	if err := validate.MaxLength("destination_label", "body", m.DestinationLabel, 50); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateDestinationType(formats strfmt.Registry) error {

	if err := validate.Required("destination_type", "body", m.DestinationType); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", *m.Slug, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", *m.Slug, 100); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", *m.Slug, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateSourceLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceLabel) { // not required
		return nil
	}

	if err := validate.MaxLength("source_label", "body", m.SourceLabel, 50); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateSourceType(formats strfmt.Registry) error {

	if err := validate.Required("source_type", "body", m.SourceType); err != nil {
		return err
	}

	return nil
}

var relationshipTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["one-to-one","one-to-many","many-to-many"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		relationshipTypeTypePropEnum = append(relationshipTypeTypePropEnum, v)
	}
}

const (

	// RelationshipTypeOneDashToDashOne captures enum value "one-to-one"
	RelationshipTypeOneDashToDashOne string = "one-to-one"

	// RelationshipTypeOneDashToDashMany captures enum value "one-to-many"
	RelationshipTypeOneDashToDashMany string = "one-to-many"

	// RelationshipTypeManyDashToDashMany captures enum value "many-to-many"
	RelationshipTypeManyDashToDashMany string = "many-to-many"
)

// prop value enum
func (m *Relationship) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, relationshipTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Relationship) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this relationship based on the context it is used
func (m *Relationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Relationship) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Relationship) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Relationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Relationship) UnmarshalBinary(b []byte) error {
	var res Relationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
