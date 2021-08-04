package models

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GoldenConfig golden config
//
// swagger:model GoldenConfig
type GoldenConfig struct {

	// custom field data
	CustomFieldData string `json:"_custom_field_data,omitempty"`

	// Backup config
	//
	// Full backup config for device.
	BackupConfig string `json:"backup_config,omitempty"`

	// Backup last attempt date
	// Format: date-time
	BackupLastAttemptDate *strfmt.DateTime `json:"backup_last_attempt_date,omitempty"`

	// Backup last success date
	// Format: date-time
	BackupLastSuccessDate *strfmt.DateTime `json:"backup_last_success_date,omitempty"`

	// Compliance config
	//
	// Full config diff for device.
	ComplianceConfig string `json:"compliance_config,omitempty"`

	// Compliance last attempt date
	// Format: date-time
	ComplianceLastAttemptDate *strfmt.DateTime `json:"compliance_last_attempt_date,omitempty"`

	// Compliance last success date
	// Format: date-time
	ComplianceLastSuccessDate *strfmt.DateTime `json:"compliance_last_success_date,omitempty"`

	// Computed fields
	// Read Only: true
	ComputedFields string `json:"computed_fields,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Device
	//
	// device
	// Required: true
	// Format: uuid
	Device *strfmt.UUID `json:"device"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Intended config
	//
	// Intended config for the device.
	IntendedConfig string `json:"intended_config,omitempty"`

	// Intended last attempt date
	// Format: date-time
	IntendedLastAttemptDate *strfmt.DateTime `json:"intended_last_attempt_date,omitempty"`

	// Intended last success date
	// Format: date-time
	IntendedLastSuccessDate *strfmt.DateTime `json:"intended_last_success_date,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this golden config
func (m *GoldenConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupLastAttemptDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupLastSuccessDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComplianceLastAttemptDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComplianceLastSuccessDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntendedLastAttemptDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntendedLastSuccessDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *GoldenConfig) validateBackupLastAttemptDate(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupLastAttemptDate) { // not required
		return nil
	}

	if err := validate.FormatOf("backup_last_attempt_date", "body", "date-time", m.BackupLastAttemptDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateBackupLastSuccessDate(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupLastSuccessDate) { // not required
		return nil
	}

	if err := validate.FormatOf("backup_last_success_date", "body", "date-time", m.BackupLastSuccessDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateComplianceLastAttemptDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ComplianceLastAttemptDate) { // not required
		return nil
	}

	if err := validate.FormatOf("compliance_last_attempt_date", "body", "date-time", m.ComplianceLastAttemptDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateComplianceLastSuccessDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ComplianceLastSuccessDate) { // not required
		return nil
	}

	if err := validate.FormatOf("compliance_last_success_date", "body", "date-time", m.ComplianceLastSuccessDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if err := validate.FormatOf("device", "body", "uuid", m.Device.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateIntendedLastAttemptDate(formats strfmt.Registry) error {
	if swag.IsZero(m.IntendedLastAttemptDate) { // not required
		return nil
	}

	if err := validate.FormatOf("intended_last_attempt_date", "body", "date-time", m.IntendedLastAttemptDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateIntendedLastSuccessDate(formats strfmt.Registry) error {
	if swag.IsZero(m.IntendedLastSuccessDate) { // not required
		return nil
	}

	if err := validate.FormatOf("intended_last_success_date", "body", "date-time", m.IntendedLastSuccessDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GoldenConfig) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this golden config based on the context it is used
func (m *GoldenConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputedFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
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

func (m *GoldenConfig) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfig) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GoldenConfig) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GoldenConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoldenConfig) UnmarshalBinary(b []byte) error {
	var res GoldenConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
