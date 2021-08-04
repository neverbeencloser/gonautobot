package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomLink custom link
//
// swagger:model CustomLink
type CustomLink struct {

	// Button class
	//
	// The class of the first link in a group will be used for the dropdown button
	// Enum: [default primary success info warning danger link]
	ButtonClass string `json:"button_class,omitempty"`

	// Content type
	// Required: true
	ContentType *string `json:"content_type"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// Group name
	//
	// Links with the same group will appear as a dropdown menu
	// Max Length: 50
	GroupName string `json:"group_name,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// New window
	//
	// Force link to open in a new window
	// Required: true
	NewWindow *bool `json:"new_window"`

	// URL
	//
	// Jinja2 template code for link URL. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.slug }}</code>.
	// Required: true
	// Max Length: 500
	// Min Length: 1
	TargetURL *string `json:"target_url"`

	// Text
	//
	// Jinja2 template code for link text. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.slug }}</code>. Links which render as empty text will not be displayed.
	// Required: true
	// Max Length: 500
	// Min Length: 1
	Text *string `json:"text"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Weight
	// Maximum: 32767
	// Minimum: 0
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this custom link
func (m *CustomLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateButtonClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewWindow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customLinkTypeButtonClassPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","primary","success","info","warning","danger","link"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customLinkTypeButtonClassPropEnum = append(customLinkTypeButtonClassPropEnum, v)
	}
}

const (

	// CustomLinkButtonClassDefault captures enum value "default"
	CustomLinkButtonClassDefault string = "default"

	// CustomLinkButtonClassPrimary captures enum value "primary"
	CustomLinkButtonClassPrimary string = "primary"

	// CustomLinkButtonClassSuccess captures enum value "success"
	CustomLinkButtonClassSuccess string = "success"

	// CustomLinkButtonClassInfo captures enum value "info"
	CustomLinkButtonClassInfo string = "info"

	// CustomLinkButtonClassWarning captures enum value "warning"
	CustomLinkButtonClassWarning string = "warning"

	// CustomLinkButtonClassDanger captures enum value "danger"
	CustomLinkButtonClassDanger string = "danger"

	// CustomLinkButtonClassLink captures enum value "link"
	CustomLinkButtonClassLink string = "link"
)

// prop value enum
func (m *CustomLink) validateButtonClassEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customLinkTypeButtonClassPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomLink) validateButtonClass(formats strfmt.Registry) error {
	if swag.IsZero(m.ButtonClass) { // not required
		return nil
	}

	// value enum
	if err := m.validateButtonClassEnum("button_class", "body", m.ButtonClass); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateGroupName(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupName) { // not required
		return nil
	}

	if err := validate.MaxLength("group_name", "body", m.GroupName, 50); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateNewWindow(formats strfmt.Registry) error {

	if err := validate.Required("new_window", "body", m.NewWindow); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateTargetURL(formats strfmt.Registry) error {

	if err := validate.Required("target_url", "body", m.TargetURL); err != nil {
		return err
	}

	if err := validate.MinLength("target_url", "body", *m.TargetURL, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("target_url", "body", *m.TargetURL, 500); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	if err := validate.MinLength("text", "body", *m.Text, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("text", "body", *m.Text, 500); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", *m.Weight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", *m.Weight, 32767, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this custom link based on the context it is used
func (m *CustomLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

func (m *CustomLink) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CustomLink) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomLink) UnmarshalBinary(b []byte) error {
	var res CustomLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
