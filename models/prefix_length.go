package models

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrefixLength prefix length
//
// swagger:model PrefixLength
type PrefixLength struct {

	// Prefix length
	// Required: true
	PrefixLength *int64 `json:"prefix_length"`
}

// Validate validates this prefix length
func (m *PrefixLength) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrefixLength(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrefixLength) validatePrefixLength(formats strfmt.Registry) error {

	if err := validate.Required("prefix_length", "body", m.PrefixLength); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this prefix length based on context it is used
func (m *PrefixLength) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrefixLength) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrefixLength) UnmarshalBinary(b []byte) error {
	var res PrefixLength
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
