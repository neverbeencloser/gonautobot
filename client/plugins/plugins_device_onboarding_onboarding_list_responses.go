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

// PluginsDeviceOnboardingOnboardingListReader is a Reader for the PluginsDeviceOnboardingOnboardingList structure.
type PluginsDeviceOnboardingOnboardingListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsDeviceOnboardingOnboardingListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsDeviceOnboardingOnboardingListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginsDeviceOnboardingOnboardingListOK creates a PluginsDeviceOnboardingOnboardingListOK with default headers values
func NewPluginsDeviceOnboardingOnboardingListOK() *PluginsDeviceOnboardingOnboardingListOK {
	return &PluginsDeviceOnboardingOnboardingListOK{}
}

/* PluginsDeviceOnboardingOnboardingListOK describes a response with status code 200, with default header values.

PluginsDeviceOnboardingOnboardingListOK plugins device onboarding onboarding list o k
*/
type PluginsDeviceOnboardingOnboardingListOK struct {
	Payload *PluginsDeviceOnboardingOnboardingListOKBody
}

func (o *PluginsDeviceOnboardingOnboardingListOK) Error() string {
	return fmt.Sprintf("[GET /plugins/device-onboarding/onboarding/][%d] pluginsDeviceOnboardingOnboardingListOK  %+v", 200, o.Payload)
}
func (o *PluginsDeviceOnboardingOnboardingListOK) GetPayload() *PluginsDeviceOnboardingOnboardingListOKBody {
	return o.Payload
}

func (o *PluginsDeviceOnboardingOnboardingListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PluginsDeviceOnboardingOnboardingListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PluginsDeviceOnboardingOnboardingListOKBody plugins device onboarding onboarding list o k body
swagger:model PluginsDeviceOnboardingOnboardingListOKBody
*/
type PluginsDeviceOnboardingOnboardingListOKBody struct {

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
	Results []*models.OnboardingTask `json:"results"`
}

// Validate validates this plugins device onboarding onboarding list o k body
func (o *PluginsDeviceOnboardingOnboardingListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PluginsDeviceOnboardingOnboardingListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("pluginsDeviceOnboardingOnboardingListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *PluginsDeviceOnboardingOnboardingListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("pluginsDeviceOnboardingOnboardingListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PluginsDeviceOnboardingOnboardingListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("pluginsDeviceOnboardingOnboardingListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PluginsDeviceOnboardingOnboardingListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("pluginsDeviceOnboardingOnboardingListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsDeviceOnboardingOnboardingListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this plugins device onboarding onboarding list o k body based on the context it is used
func (o *PluginsDeviceOnboardingOnboardingListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PluginsDeviceOnboardingOnboardingListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginsDeviceOnboardingOnboardingListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PluginsDeviceOnboardingOnboardingListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PluginsDeviceOnboardingOnboardingListOKBody) UnmarshalBinary(b []byte) error {
	var res PluginsDeviceOnboardingOnboardingListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
