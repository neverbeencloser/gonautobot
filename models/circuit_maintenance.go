package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CircuitMaintenance circuit maintenance
//
// swagger:model CircuitMaintenance
type CircuitMaintenance struct {

	// Ack
	Ack *bool `json:"ack,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// End time
	// Required: true
	// Format: date-time
	EndTime *strfmt.DateTime `json:"end_time"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name
	// Max Length: 100
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// Start time
	// Required: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time"`

	// Status
	// Enum: [TENTATIVE CONFIRMED CANCELLED IN-PROCESS COMPLETED RE-SCHEDULED]
	Status *string `json:"status,omitempty"`
}

// Validate validates this circuit maintenance
func (m *CircuitMaintenance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CircuitMaintenance) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("end_time", "body", m.EndTime); err != nil {
		return err
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CircuitMaintenance) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CircuitMaintenance) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *CircuitMaintenance) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("start_time", "body", m.StartTime); err != nil {
		return err
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var circuitMaintenanceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TENTATIVE","CONFIRMED","CANCELLED","IN-PROCESS","COMPLETED","RE-SCHEDULED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		circuitMaintenanceTypeStatusPropEnum = append(circuitMaintenanceTypeStatusPropEnum, v)
	}
}

const (

	// CircuitMaintenanceStatusTENTATIVE captures enum value "TENTATIVE"
	CircuitMaintenanceStatusTENTATIVE string = "TENTATIVE"

	// CircuitMaintenanceStatusCONFIRMED captures enum value "CONFIRMED"
	CircuitMaintenanceStatusCONFIRMED string = "CONFIRMED"

	// CircuitMaintenanceStatusCANCELLED captures enum value "CANCELLED"
	CircuitMaintenanceStatusCANCELLED string = "CANCELLED"

	// CircuitMaintenanceStatusINDashPROCESS captures enum value "IN-PROCESS"
	CircuitMaintenanceStatusINDashPROCESS string = "IN-PROCESS"

	// CircuitMaintenanceStatusCOMPLETED captures enum value "COMPLETED"
	CircuitMaintenanceStatusCOMPLETED string = "COMPLETED"

	// CircuitMaintenanceStatusREDashSCHEDULED captures enum value "RE-SCHEDULED"
	CircuitMaintenanceStatusREDashSCHEDULED string = "RE-SCHEDULED"
)

// prop value enum
func (m *CircuitMaintenance) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, circuitMaintenanceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CircuitMaintenance) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this circuit maintenance based on the context it is used
func (m *CircuitMaintenance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CircuitMaintenance) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CircuitMaintenance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CircuitMaintenance) UnmarshalBinary(b []byte) error {
	var res CircuitMaintenance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
