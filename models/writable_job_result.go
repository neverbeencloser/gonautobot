package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableJobResult writable job result
//
// swagger:model WritableJobResult
type WritableJobResult struct {

	// Completed
	// Format: date-time
	Completed *strfmt.DateTime `json:"completed,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Data
	Data *string `json:"data,omitempty"`

	// Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Job id
	// Required: true
	// Format: uuid
	JobID *strfmt.UUID `json:"job_id"`

	// Name
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Name *string `json:"name"`

	// Obj type
	// Read Only: true
	ObjType string `json:"obj_type,omitempty"`

	// Status
	// Enum: [pending running completed errored failed]
	Status string `json:"status,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// User
	// Format: uuid
	User *strfmt.UUID `json:"user,omitempty"`
}

// Validate validates this writable job result
func (m *WritableJobResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableJobResult) validateCompleted(formats strfmt.Registry) error {
	if swag.IsZero(m.Completed) { // not required
		return nil
	}

	if err := validate.FormatOf("completed", "body", "date-time", m.Completed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableJobResult) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableJobResult) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableJobResult) validateJobID(formats strfmt.Registry) error {

	if err := validate.Required("job_id", "body", m.JobID); err != nil {
		return err
	}

	if err := validate.FormatOf("job_id", "body", "uuid", m.JobID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableJobResult) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 255); err != nil {
		return err
	}

	return nil
}

var writableJobResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","running","completed","errored","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableJobResultTypeStatusPropEnum = append(writableJobResultTypeStatusPropEnum, v)
	}
}

const (

	// WritableJobResultStatusPending captures enum value "pending"
	WritableJobResultStatusPending string = "pending"

	// WritableJobResultStatusRunning captures enum value "running"
	WritableJobResultStatusRunning string = "running"

	// WritableJobResultStatusCompleted captures enum value "completed"
	WritableJobResultStatusCompleted string = "completed"

	// WritableJobResultStatusErrored captures enum value "errored"
	WritableJobResultStatusErrored string = "errored"

	// WritableJobResultStatusFailed captures enum value "failed"
	WritableJobResultStatusFailed string = "failed"
)

// prop value enum
func (m *WritableJobResult) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableJobResultTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableJobResult) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableJobResult) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableJobResult) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if err := validate.FormatOf("user", "body", "uuid", m.User.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable job result based on the context it is used
func (m *WritableJobResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjType(ctx, formats); err != nil {
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

func (m *WritableJobResult) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableJobResult) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableJobResult) contextValidateObjType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "obj_type", "body", string(m.ObjType)); err != nil {
		return err
	}

	return nil
}

func (m *WritableJobResult) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableJobResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableJobResult) UnmarshalBinary(b []byte) error {
	var res WritableJobResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
