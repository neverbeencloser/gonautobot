package models

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExportTemplate export template
//
// swagger:model ExportTemplate
type ExportTemplate struct {

	// Content type
	// Required: true
	ContentType *string `json:"content_type"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	//
	// Human friendly display value
	// Read Only: true
	// Min Length: 1
	Display string `json:"display,omitempty"`

	// File extension
	//
	// Extension to append to the rendered filename
	// Max Length: 15
	FileExtension string `json:"file_extension,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// MIME type
	//
	// Defaults to <code>text/plain</code>
	// Max Length: 50
	MimeType string `json:"mime_type,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Owner
	// Read Only: true
	Owner map[string]*string `json:"owner,omitempty"`

	// Owner content type
	OwnerContentType *string `json:"owner_content_type,omitempty"`

	// Owner object id
	// Format: uuid
	OwnerObjectID *strfmt.UUID `json:"owner_object_id,omitempty"`

	// Template code
	//
	// The list of objects being exported is passed as a context variable named <code>queryset</code>.
	// Required: true
	// Min Length: 1
	TemplateCode *string `json:"template_code"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this export template
func (m *ExportTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileExtension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMimeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplateCode(formats); err != nil {
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

func (m *ExportTemplate) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateFileExtension(formats strfmt.Registry) error {
	if swag.IsZero(m.FileExtension) { // not required
		return nil
	}

	if err := validate.MaxLength("file_extension", "body", m.FileExtension, 15); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateMimeType(formats strfmt.Registry) error {
	if swag.IsZero(m.MimeType) { // not required
		return nil
	}

	if err := validate.MaxLength("mime_type", "body", m.MimeType, 50); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateName(formats strfmt.Registry) error {

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

func (m *ExportTemplate) validateOwnerObjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerObjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("owner_object_id", "body", "uuid", m.OwnerObjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateTemplateCode(formats strfmt.Registry) error {

	if err := validate.Required("template_code", "body", m.TemplateCode); err != nil {
		return err
	}

	if err := validate.MinLength("template_code", "body", *m.TemplateCode, 1); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this export template based on the context it is used
func (m *ExportTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
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

func (m *ExportTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ExportTemplate) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ExportTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExportTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportTemplate) UnmarshalBinary(b []byte) error {
	var res ExportTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
