package models

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableRelationshipAssociation writable relationship association
//
// swagger:model WritableRelationshipAssociation
type WritableRelationshipAssociation struct {

	// Destination id
	// Required: true
	// Format: uuid
	DestinationID *strfmt.UUID `json:"destination_id"`

	// Destination type
	// Required: true
	DestinationType *string `json:"destination_type"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Relationship
	// Required: true
	// Format: uuid
	Relationship *strfmt.UUID `json:"relationship"`

	// Source id
	// Required: true
	// Format: uuid
	SourceID *strfmt.UUID `json:"source_id"`

	// Source type
	// Required: true
	SourceType *string `json:"source_type"`
}

// Validate validates this writable relationship association
func (m *WritableRelationshipAssociation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRelationshipAssociation) validateDestinationID(formats strfmt.Registry) error {

	if err := validate.Required("destination_id", "body", m.DestinationID); err != nil {
		return err
	}

	if err := validate.FormatOf("destination_id", "body", "uuid", m.DestinationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRelationshipAssociation) validateDestinationType(formats strfmt.Registry) error {

	if err := validate.Required("destination_type", "body", m.DestinationType); err != nil {
		return err
	}

	return nil
}

func (m *WritableRelationshipAssociation) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRelationshipAssociation) validateRelationship(formats strfmt.Registry) error {

	if err := validate.Required("relationship", "body", m.Relationship); err != nil {
		return err
	}

	if err := validate.FormatOf("relationship", "body", "uuid", m.Relationship.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRelationshipAssociation) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("source_id", "body", m.SourceID); err != nil {
		return err
	}

	if err := validate.FormatOf("source_id", "body", "uuid", m.SourceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableRelationshipAssociation) validateSourceType(formats strfmt.Registry) error {

	if err := validate.Required("source_type", "body", m.SourceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable relationship association based on the context it is used
func (m *WritableRelationshipAssociation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRelationshipAssociation) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableRelationshipAssociation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableRelationshipAssociation) UnmarshalBinary(b []byte) error {
	var res WritableRelationshipAssociation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
