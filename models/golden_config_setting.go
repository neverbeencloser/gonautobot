package models

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GoldenConfigSetting golden config setting
//
// swagger:model GoldenConfigSetting
type GoldenConfigSetting struct {

	// custom field data
	CustomFieldData string `json:"_custom_field_data,omitempty"`

	// Backup Path in Jinja Template Form
	//
	// The Jinja path representation of where the backup file will be found. The variable `obj` is available as the device instance object of a given device, as is the case for all Jinja templates. e.g. `{{obj.site.slug}}/{{obj.name}}.cfg`
	// Max Length: 255
	BackupPathTemplate string `json:"backup_path_template,omitempty"`

	// Backup repository
	// Format: uuid
	BackupRepository *strfmt.UUID `json:"backup_repository,omitempty"`

	// Backup Test
	//
	// Whether or not to pretest the connectivity of the device by verifying there is a resolvable IP that can connect to port 22.
	BackupTestConnectivity bool `json:"backup_test_connectivity,omitempty"`

	// Computed fields
	// Read Only: true
	ComputedFields string `json:"computed_fields,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

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

	// Intended Path in Jinja Template Form
	//
	// The Jinja path representation of where the generated file will be places. e.g. `{{obj.site.slug}}/{{obj.name}}.cfg`
	// Max Length: 255
	IntendedPathTemplate string `json:"intended_path_template,omitempty"`

	// Intended repository
	// Format: uuid
	IntendedRepository *strfmt.UUID `json:"intended_repository,omitempty"`

	// Template Path in Jinja Template Form
	//
	// The Jinja path representation of where the Jinja template can be found. e.g. `{{obj.platform.slug}}.j2`
	// Max Length: 255
	JinjaPathTemplate string `json:"jinja_path_template,omitempty"`

	// Jinja repository
	// Format: uuid
	JinjaRepository *strfmt.UUID `json:"jinja_repository,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Scope
	//
	// Queryset filter matching the list of devices for the scope of devices to be considered.
	Scope *string `json:"scope,omitempty"`

	// GraphQL Query
	//
	// A query starting with `query ($device_id: ID!)` that is used to render the config. Please make sure to alias name, see FAQ for more details.
	SotAggQuery string `json:"sot_agg_query,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this golden config setting
func (m *GoldenConfigSetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupPathTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntendedPathTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntendedRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJinjaPathTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJinjaRepository(formats); err != nil {
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

func (m *GoldenConfigSetting) validateBackupPathTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupPathTemplate) { // not required
		return nil
	}

	if err := validate.MaxLength("backup_path_template", "body", m.BackupPathTemplate, 255); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) validateBackupRepository(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupRepository) { // not required
		return nil
	}

	if err := validate.FormatOf("backup_repository", "body", "uuid", m.BackupRepository.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) validateIntendedPathTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.IntendedPathTemplate) { // not required
		return nil
	}

	if err := validate.MaxLength("intended_path_template", "body", m.IntendedPathTemplate, 255); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) validateIntendedRepository(formats strfmt.Registry) error {
	if swag.IsZero(m.IntendedRepository) { // not required
		return nil
	}

	if err := validate.FormatOf("intended_repository", "body", "uuid", m.IntendedRepository.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) validateJinjaPathTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.JinjaPathTemplate) { // not required
		return nil
	}

	if err := validate.MaxLength("jinja_path_template", "body", m.JinjaPathTemplate, 255); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) validateJinjaRepository(formats strfmt.Registry) error {
	if swag.IsZero(m.JinjaRepository) { // not required
		return nil
	}

	if err := validate.FormatOf("jinja_repository", "body", "uuid", m.JinjaRepository.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) validateTags(formats strfmt.Registry) error {
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

func (m *GoldenConfigSetting) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this golden config setting based on the context it is used
func (m *GoldenConfigSetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GoldenConfigSetting) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *GoldenConfigSetting) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GoldenConfigSetting) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GoldenConfigSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoldenConfigSetting) UnmarshalBinary(b []byte) error {
	var res GoldenConfigSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
