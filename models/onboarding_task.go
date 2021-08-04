package models

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OnboardingTask onboarding task
//
// swagger:model OnboardingTask
type OnboardingTask struct {

	// Created device
	//
	// Created device name
	// Read Only: true
	// Min Length: 1
	CreatedDevice string `json:"created_device,omitempty"`

	// Device type
	//
	// Nautobot device type 'slug' value
	// Min Length: 1
	DeviceType string `json:"device_type,omitempty"`

	// Failed reason
	//
	// Failure reason
	// Read Only: true
	// Min Length: 1
	FailedReason string `json:"failed_reason,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Ip address
	//
	// IP Address to reach device
	// Required: true
	// Min Length: 1
	IPAddress *string `json:"ip_address"`

	// Message
	//
	// Status message
	// Read Only: true
	// Min Length: 1
	Message string `json:"message,omitempty"`

	// Password
	//
	// Device password
	// Min Length: 1
	Password string `json:"password,omitempty"`

	// Platform
	//
	// Nautobot Platform 'slug' value
	// Pattern: ^[-a-zA-Z0-9_]+$
	Platform string `json:"platform,omitempty"`

	// Port
	//
	// Device PORT to check for online
	Port int64 `json:"port,omitempty"`

	// Role
	//
	// Nautobot device role 'slug' value
	// Pattern: ^[-a-zA-Z0-9_]+$
	Role string `json:"role,omitempty"`

	// Secret
	//
	// Device secret password
	// Min Length: 1
	Secret string `json:"secret,omitempty"`

	// Site
	//
	// Nautobot site 'slug' value
	// Required: true
	// Pattern: ^[-a-zA-Z0-9_]+$
	Site *string `json:"site"`

	// Status
	//
	// Onboarding Status
	// Read Only: true
	// Min Length: 1
	Status string `json:"status,omitempty"`

	// Timeout
	//
	// Timeout (sec) for device connect
	Timeout int64 `json:"timeout,omitempty"`

	// Username
	//
	// Device username
	// Min Length: 1
	Username string `json:"username,omitempty"`
}

// Validate validates this onboarding task
func (m *OnboardingTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailedReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnboardingTask) validateCreatedDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDevice) { // not required
		return nil
	}

	if err := validate.MinLength("created_device", "body", m.CreatedDevice, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validateDeviceType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceType) { // not required
		return nil
	}

	if err := validate.MinLength("device_type", "body", m.DeviceType, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validateFailedReason(formats strfmt.Registry) error {
	if swag.IsZero(m.FailedReason) { // not required
		return nil
	}

	if err := validate.MinLength("failed_reason", "body", m.FailedReason, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ip_address", "body", m.IPAddress); err != nil {
		return err
	}

	if err := validate.MinLength("ip_address", "body", *m.IPAddress, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if err := validate.MinLength("message", "body", m.Message, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("password", "body", m.Password, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if err := validate.Pattern("platform", "body", m.Platform, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if err := validate.Pattern("role", "body", m.Role, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validateSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.Secret) { // not required
		return nil
	}

	if err := validate.MinLength("secret", "body", m.Secret, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	if err := validate.Pattern("site", "body", *m.Site, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.MinLength("status", "body", m.Status, 1); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) validateUsername(formats strfmt.Registry) error {
	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := validate.MinLength("username", "body", m.Username, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this onboarding task based on the context it is used
func (m *OnboardingTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFailedReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnboardingTask) contextValidateCreatedDevice(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_device", "body", string(m.CreatedDevice)); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) contextValidateFailedReason(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "failed_reason", "body", string(m.FailedReason)); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

func (m *OnboardingTask) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnboardingTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnboardingTask) UnmarshalBinary(b []byte) error {
	var res OnboardingTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
