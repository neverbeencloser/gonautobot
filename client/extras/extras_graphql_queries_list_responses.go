package extras

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

// ExtrasGraphqlQueriesListReader is a Reader for the ExtrasGraphqlQueriesList structure.
type ExtrasGraphqlQueriesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasGraphqlQueriesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasGraphqlQueriesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasGraphqlQueriesListOK creates a ExtrasGraphqlQueriesListOK with default headers values
func NewExtrasGraphqlQueriesListOK() *ExtrasGraphqlQueriesListOK {
	return &ExtrasGraphqlQueriesListOK{}
}

/* ExtrasGraphqlQueriesListOK describes a response with status code 200, with default header values.

ExtrasGraphqlQueriesListOK extras graphql queries list o k
*/
type ExtrasGraphqlQueriesListOK struct {
	Payload *ExtrasGraphqlQueriesListOKBody
}

func (o *ExtrasGraphqlQueriesListOK) Error() string {
	return fmt.Sprintf("[GET /extras/graphql-queries/][%d] extrasGraphqlQueriesListOK  %+v", 200, o.Payload)
}
func (o *ExtrasGraphqlQueriesListOK) GetPayload() *ExtrasGraphqlQueriesListOKBody {
	return o.Payload
}

func (o *ExtrasGraphqlQueriesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ExtrasGraphqlQueriesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ExtrasGraphqlQueriesListOKBody extras graphql queries list o k body
swagger:model ExtrasGraphqlQueriesListOKBody
*/
type ExtrasGraphqlQueriesListOKBody struct {

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
	Results []*models.GraphQLQuery `json:"results"`
}

// Validate validates this extras graphql queries list o k body
func (o *ExtrasGraphqlQueriesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ExtrasGraphqlQueriesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("extrasGraphqlQueriesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasGraphqlQueriesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasGraphqlQueriesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasGraphqlQueriesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("extrasGraphqlQueriesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ExtrasGraphqlQueriesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("extrasGraphqlQueriesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasGraphqlQueriesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this extras graphql queries list o k body based on the context it is used
func (o *ExtrasGraphqlQueriesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ExtrasGraphqlQueriesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extrasGraphqlQueriesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ExtrasGraphqlQueriesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ExtrasGraphqlQueriesListOKBody) UnmarshalBinary(b []byte) error {
	var res ExtrasGraphqlQueriesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
