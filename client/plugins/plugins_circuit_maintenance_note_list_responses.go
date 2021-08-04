package plugins

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// PluginsCircuitMaintenanceNoteListReader is a Reader for the PluginsCircuitMaintenanceNoteList structure.
type PluginsCircuitMaintenanceNoteListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsCircuitMaintenanceNoteListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsCircuitMaintenanceNoteListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsCircuitMaintenanceNoteListOK creates a PluginsCircuitMaintenanceNoteListOK with default headers values
func NewPluginsCircuitMaintenanceNoteListOK() *PluginsCircuitMaintenanceNoteListOK {
	return &PluginsCircuitMaintenanceNoteListOK{}
}

/* PluginsCircuitMaintenanceNoteListOK describes a response with status code 200, with default header values.

PluginsCircuitMaintenanceNoteListOK plugins circuit maintenance note list o k
*/
type PluginsCircuitMaintenanceNoteListOK struct {
	Payload *PluginsCircuitMaintenanceNoteListOKBody
}

func (o *PluginsCircuitMaintenanceNoteListOK) Error() string {
	return fmt.Sprintf("[GET /plugins/circuit-maintenance/note/][%d] pluginsCircuitMaintenanceNoteListOK  %+v", 200, o.Payload)
}
func (o *PluginsCircuitMaintenanceNoteListOK) GetPayload() *PluginsCircuitMaintenanceNoteListOKBody {
	return o.Payload
}

func (o *PluginsCircuitMaintenanceNoteListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PluginsCircuitMaintenanceNoteListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PluginsCircuitMaintenanceNoteListOKBody plugins circuit maintenance note list o k body
swagger:model PluginsCircuitMaintenanceNoteListOKBody
*/
type PluginsCircuitMaintenanceNoteListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.Note `json:"results"`
}

// Validate validates this plugins circuit maintenance note list o k body
func (o *PluginsCircuitMaintenanceNoteListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PluginsCircuitMaintenanceNoteListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("pluginsCircuitMaintenanceNoteListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *PluginsCircuitMaintenanceNoteListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("pluginsCircuitMaintenanceNoteListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PluginsCircuitMaintenanceNoteListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("pluginsCircuitMaintenanceNoteListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PluginsCircuitMaintenanceNoteListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("pluginsCircuitMaintenanceNoteListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsCircuitMaintenanceNoteListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plugins circuit maintenance note list o k body based on the context it is used
func (o *PluginsCircuitMaintenanceNoteListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PluginsCircuitMaintenanceNoteListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsCircuitMaintenanceNoteListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PluginsCircuitMaintenanceNoteListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PluginsCircuitMaintenanceNoteListOKBody) UnmarshalBinary(b []byte) error {
	var res PluginsCircuitMaintenanceNoteListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
