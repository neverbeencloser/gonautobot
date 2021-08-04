package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableFrontPortTemplate writable front port template
//
// swagger:model WritableFrontPortTemplate
type WritableFrontPortTemplate struct {

	// Computed fields
	// Read Only: true
	ComputedFields string `json:"computed_fields,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device type
	// Required: true
	// Format: uuid
	DeviceType *strfmt.UUID `json:"device_type"`

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

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Rear port
	// Required: true
	// Format: uuid
	RearPort *strfmt.UUID `json:"rear_port"`

	// Rear port position
	// Maximum: 1024
	// Minimum: 1
	RearPortPosition int64 `json:"rear_port_position,omitempty"`

	// Type
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st cs sn splice]
	Type *string `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable front port template
func (m *WritableFrontPortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPortPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *WritableFrontPortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if err := validate.FormatOf("device_type", "body", "uuid", m.DeviceType.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateRearPort(formats strfmt.Registry) error {

	if err := validate.Required("rear_port", "body", m.RearPort); err != nil {
		return err
	}

	if err := validate.FormatOf("rear_port", "body", "uuid", m.RearPort.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateRearPortPosition(formats strfmt.Registry) error {
	if swag.IsZero(m.RearPortPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("rear_port_position", "body", m.RearPortPosition, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rear_port_position", "body", m.RearPortPosition, 1024, false); err != nil {
		return err
	}

	return nil
}

var writableFrontPortTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st","cs","sn","splice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableFrontPortTemplateTypeTypePropEnum = append(writableFrontPortTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritableFrontPortTemplateTypeNr8p8c captures enum value "8p8c"
	WritableFrontPortTemplateTypeNr8p8c string = "8p8c"

	// WritableFrontPortTemplateTypeNr8p6c captures enum value "8p6c"
	WritableFrontPortTemplateTypeNr8p6c string = "8p6c"

	// WritableFrontPortTemplateTypeNr8p4c captures enum value "8p4c"
	WritableFrontPortTemplateTypeNr8p4c string = "8p4c"

	// WritableFrontPortTemplateTypeNr8p2c captures enum value "8p2c"
	WritableFrontPortTemplateTypeNr8p2c string = "8p2c"

	// WritableFrontPortTemplateTypeGg45 captures enum value "gg45"
	WritableFrontPortTemplateTypeGg45 string = "gg45"

	// WritableFrontPortTemplateTypeTeraDash4p captures enum value "tera-4p"
	WritableFrontPortTemplateTypeTeraDash4p string = "tera-4p"

	// WritableFrontPortTemplateTypeTeraDash2p captures enum value "tera-2p"
	WritableFrontPortTemplateTypeTeraDash2p string = "tera-2p"

	// WritableFrontPortTemplateTypeTeraDash1p captures enum value "tera-1p"
	WritableFrontPortTemplateTypeTeraDash1p string = "tera-1p"

	// WritableFrontPortTemplateTypeNr110DashPunch captures enum value "110-punch"
	WritableFrontPortTemplateTypeNr110DashPunch string = "110-punch"

	// WritableFrontPortTemplateTypeBnc captures enum value "bnc"
	WritableFrontPortTemplateTypeBnc string = "bnc"

	// WritableFrontPortTemplateTypeMrj21 captures enum value "mrj21"
	WritableFrontPortTemplateTypeMrj21 string = "mrj21"

	// WritableFrontPortTemplateTypeFc captures enum value "fc"
	WritableFrontPortTemplateTypeFc string = "fc"

	// WritableFrontPortTemplateTypeLc captures enum value "lc"
	WritableFrontPortTemplateTypeLc string = "lc"

	// WritableFrontPortTemplateTypeLcDashApc captures enum value "lc-apc"
	WritableFrontPortTemplateTypeLcDashApc string = "lc-apc"

	// WritableFrontPortTemplateTypeLsh captures enum value "lsh"
	WritableFrontPortTemplateTypeLsh string = "lsh"

	// WritableFrontPortTemplateTypeLshDashApc captures enum value "lsh-apc"
	WritableFrontPortTemplateTypeLshDashApc string = "lsh-apc"

	// WritableFrontPortTemplateTypeMpo captures enum value "mpo"
	WritableFrontPortTemplateTypeMpo string = "mpo"

	// WritableFrontPortTemplateTypeMtrj captures enum value "mtrj"
	WritableFrontPortTemplateTypeMtrj string = "mtrj"

	// WritableFrontPortTemplateTypeSc captures enum value "sc"
	WritableFrontPortTemplateTypeSc string = "sc"

	// WritableFrontPortTemplateTypeScDashApc captures enum value "sc-apc"
	WritableFrontPortTemplateTypeScDashApc string = "sc-apc"

	// WritableFrontPortTemplateTypeSt captures enum value "st"
	WritableFrontPortTemplateTypeSt string = "st"

	// WritableFrontPortTemplateTypeCs captures enum value "cs"
	WritableFrontPortTemplateTypeCs string = "cs"

	// WritableFrontPortTemplateTypeSn captures enum value "sn"
	WritableFrontPortTemplateTypeSn string = "sn"

	// WritableFrontPortTemplateTypeSplice captures enum value "splice"
	WritableFrontPortTemplateTypeSplice string = "splice"
)

// prop value enum
func (m *WritableFrontPortTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableFrontPortTemplateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableFrontPortTemplate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable front port template based on the context it is used
func (m *WritableFrontPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputedFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
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

func (m *WritableFrontPortTemplate) contextValidateComputedFields(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computed_fields", "body", string(m.ComputedFields)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableFrontPortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableFrontPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableFrontPortTemplate) UnmarshalBinary(b []byte) error {
	var res WritableFrontPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
