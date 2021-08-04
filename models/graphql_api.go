package models

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GraphQLAPI graph q l API
//
// swagger:model GraphQLAPI
type GraphQLAPI struct {

	// Query
	//
	// GraphQL query
	// Required: true
	// Min Length: 1
	Query *string `json:"query"`

	// Variables
	//
	// Variables in JSON Format
	Variables string `json:"variables,omitempty"`
}

// Validate validates this graph q l API
func (m *GraphQLAPI) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphQLAPI) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	if err := validate.MinLength("query", "body", *m.Query, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this graph q l API based on context it is used
func (m *GraphQLAPI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GraphQLAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphQLAPI) UnmarshalBinary(b []byte) error {
	var res GraphQLAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
