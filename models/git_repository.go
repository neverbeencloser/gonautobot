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

// GitRepository git repository
//
// swagger:model GitRepository
type GitRepository struct {

	// Branch
	// Max Length: 64
	// Min Length: 1
	Branch string `json:"branch,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Current head
	//
	// Commit hash of the most recent fetch from the selected branch. Used for syncing between workers.
	// Max Length: 48
	CurrentHead string `json:"current_head,omitempty"`

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

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// provided contents
	ProvidedContents []string `json:"provided_contents"`

	// Remote url
	//
	// Only HTTP and HTTPS URLs are presently supported
	// Required: true
	// Max Length: 255
	// Min Length: 1
	// Format: uri
	RemoteURL *strfmt.URI `json:"remote_url"`

	// Slug
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// Token
	// Min Length: 1
	Token string `json:"token,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Username
	// Max Length: 64
	Username string `json:"username,omitempty"`
}

// Validate validates this git repository
func (m *GitRepository) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentHead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvidedContents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
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

func (m *GitRepository) validateBranch(formats strfmt.Registry) error {
	if swag.IsZero(m.Branch) { // not required
		return nil
	}

	if err := validate.MinLength("branch", "body", m.Branch, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("branch", "body", m.Branch, 64); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) validateCurrentHead(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentHead) { // not required
		return nil
	}

	if err := validate.MaxLength("current_head", "body", m.CurrentHead, 48); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) validateDisplay(formats strfmt.Registry) error {
	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if err := validate.MinLength("display", "body", m.Display, 1); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) validateName(formats strfmt.Registry) error {

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

var gitRepositoryProvidedContentsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["extras.configcontext","extras.configcontextschema","extras.exporttemplate","extras.job"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitRepositoryProvidedContentsItemsEnum = append(gitRepositoryProvidedContentsItemsEnum, v)
	}
}

func (m *GitRepository) validateProvidedContentsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gitRepositoryProvidedContentsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GitRepository) validateProvidedContents(formats strfmt.Registry) error {
	if swag.IsZero(m.ProvidedContents) { // not required
		return nil
	}

	for i := 0; i < len(m.ProvidedContents); i++ {

		// value enum
		if err := m.validateProvidedContentsItemsEnum("provided_contents"+"."+strconv.Itoa(i), "body", m.ProvidedContents[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GitRepository) validateRemoteURL(formats strfmt.Registry) error {

	if err := validate.Required("remote_url", "body", m.RemoteURL); err != nil {
		return err
	}

	if err := validate.MinLength("remote_url", "body", m.RemoteURL.String(), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("remote_url", "body", m.RemoteURL.String(), 255); err != nil {
		return err
	}

	if err := validate.FormatOf("remote_url", "body", "uri", m.RemoteURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", *m.Slug, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", *m.Slug, 100); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", *m.Slug, `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) validateToken(formats strfmt.Registry) error {
	if swag.IsZero(m.Token) { // not required
		return nil
	}

	if err := validate.MinLength("token", "body", m.Token, 1); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) validateUsername(formats strfmt.Registry) error {
	if swag.IsZero(m.Username) { // not required
		return nil
	}

	if err := validate.MaxLength("username", "body", m.Username, 64); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this git repository based on the context it is used
func (m *GitRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitRepository) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *GitRepository) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitRepository) UnmarshalBinary(b []byte) error {
	var res GitRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
