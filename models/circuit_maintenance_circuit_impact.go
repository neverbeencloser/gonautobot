package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CircuitMaintenanceCircuitImpact circuit maintenance circuit impact
//
// swagger:model CircuitMaintenanceCircuitImpact
type CircuitMaintenanceCircuitImpact struct {

	// Circuit
	// Required: true
	// Format: uuid
	Circuit *strfmt.UUID `json:"circuit"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Impact
	// Enum: [NO-IMPACT REDUCED-REDUNDANCY DEGRADED OUTAGE]
	Impact *string `json:"impact,omitempty"`

	// Maintenance
	// Required: true
	// Format: uuid
	Maintenance *strfmt.UUID `json:"maintenance"`
}

// Validate validates this circuit maintenance circuit impact
func (m *CircuitMaintenanceCircuitImpact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCircuit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CircuitMaintenanceCircuitImpact) validateCircuit(formats strfmt.Registry) error {

	if err := validate.Required("circuit", "body", m.Circuit); err != nil {
		return err
	}

	if err := validate.FormatOf("circuit", "body", "uuid", m.Circuit.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CircuitMaintenanceCircuitImpact) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var circuitMaintenanceCircuitImpactTypeImpactPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NO-IMPACT","REDUCED-REDUNDANCY","DEGRADED","OUTAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		circuitMaintenanceCircuitImpactTypeImpactPropEnum = append(circuitMaintenanceCircuitImpactTypeImpactPropEnum, v)
	}
}

const (

	// CircuitMaintenanceCircuitImpactImpactNODashIMPACT captures enum value "NO-IMPACT"
	CircuitMaintenanceCircuitImpactImpactNODashIMPACT string = "NO-IMPACT"

	// CircuitMaintenanceCircuitImpactImpactREDUCEDDashREDUNDANCY captures enum value "REDUCED-REDUNDANCY"
	CircuitMaintenanceCircuitImpactImpactREDUCEDDashREDUNDANCY string = "REDUCED-REDUNDANCY"

	// CircuitMaintenanceCircuitImpactImpactDEGRADED captures enum value "DEGRADED"
	CircuitMaintenanceCircuitImpactImpactDEGRADED string = "DEGRADED"

	// CircuitMaintenanceCircuitImpactImpactOUTAGE captures enum value "OUTAGE"
	CircuitMaintenanceCircuitImpactImpactOUTAGE string = "OUTAGE"
)

// prop value enum
func (m *CircuitMaintenanceCircuitImpact) validateImpactEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, circuitMaintenanceCircuitImpactTypeImpactPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CircuitMaintenanceCircuitImpact) validateImpact(formats strfmt.Registry) error {
	if swag.IsZero(m.Impact) { // not required
		return nil
	}

	// value enum
	if err := m.validateImpactEnum("impact", "body", *m.Impact); err != nil {
		return err
	}

	return nil
}

func (m *CircuitMaintenanceCircuitImpact) validateMaintenance(formats strfmt.Registry) error {

	if err := validate.Required("maintenance", "body", m.Maintenance); err != nil {
		return err
	}

	if err := validate.FormatOf("maintenance", "body", "uuid", m.Maintenance.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this circuit maintenance circuit impact based on the context it is used
func (m *CircuitMaintenanceCircuitImpact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CircuitMaintenanceCircuitImpact) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CircuitMaintenanceCircuitImpact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CircuitMaintenanceCircuitImpact) UnmarshalBinary(b []byte) error {
	var res CircuitMaintenanceCircuitImpact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
