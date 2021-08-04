package models

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ComplianceRule compliance rule
//
// swagger:model ComplianceRule
type ComplianceRule struct {

	// custom field data
	CustomFieldData string `json:"_custom_field_data,omitempty"`

	// Computed fields
	// Read Only: true
	ComputedFields string `json:"computed_fields,omitempty"`

	// Configured Ordered
	//
	// Whether or not the configuration order matters, such as in ACLs.
	// Required: true
	ConfigOrdered *bool `json:"config_ordered"`

	// Config type
	//
	// Whether the config is in cli or json/structured format.
	// Enum: [cli json]
	ConfigType string `json:"config_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// Feature
	// Required: true
	// Format: uuid
	Feature *strfmt.UUID `json:"feature"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Config to Match
	//
	// The config to match that is matched based on the parent most configuration. e.g. `router bgp` or `ntp`.
	// Min Length: 1
	MatchConfig *string `json:"match_config,omitempty"`

	// Platform
	// Required: true
	// Format: uuid
	Platform *strfmt.UUID `json:"platform"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this compliance rule
func (m *ComplianceRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigOrdered(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
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

func (m *ComplianceRule) validateConfigOrdered(formats strfmt.Registry) error {

	if err := validate.Required("config_ordered", "body", m.ConfigOrdered); err != nil {
		return err
	}

	return nil
}

var complianceRuleTypeConfigTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cli","json"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		complianceRuleTypeConfigTypePropEnum = append(complianceRuleTypeConfigTypePropEnum, v)
	}
}

const (

	// ComplianceRuleConfigTypeCli captures enum value "cli"
	ComplianceRuleConfigTypeCli string = "cli"

	// ComplianceRuleConfigTypeJSON captures enum value "json"
	ComplianceRuleConfigTypeJSON string = "json"
)

// prop value enum
func (m *ComplianceRule) validateConfigTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, complianceRuleTypeConfigTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComplianceRule) validateConfigType(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigType) { // not required
		return nil
	}

	// value enum
	if err := m.validateConfigTypeEnum("config_type", "body", m.ConfigType); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	if err := validate.FormatOf("feature", "body", "uuid", m.Feature.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) validateMatchConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchConfig) { // not required
		return nil
	}

	if err := validate.MinLength("match_config", "body", *m.MatchConfig, 1); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	if err := validate.FormatOf("platform", "body", "uuid", m.Platform.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) validateTags(formats strfmt.Registry) error {
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

func (m *ComplianceRule) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this compliance rule based on the context it is used
func (m *ComplianceRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *ComplianceRule) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceRule) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ComplianceRule) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComplianceRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComplianceRule) UnmarshalBinary(b []byte) error {
	var res ComplianceRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
