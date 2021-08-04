package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobInput job input
//
// swagger:model JobInput
type JobInput struct {

	// Commit
	Commit bool `json:"commit,omitempty"`

	// Data
	Data string `json:"data,omitempty"`
}

// Validate validates this job input
func (m *JobInput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this job input based on context it is used
func (m *JobInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobInput) UnmarshalBinary(b []byte) error {
	var res JobInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
