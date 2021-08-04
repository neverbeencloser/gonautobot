package graphql

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GraphqlCreateReader is a Reader for the GraphqlCreate structure.
type GraphqlCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GraphqlCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGraphqlCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGraphqlCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGraphqlCreateOK creates a GraphqlCreateOK with default headers values
func NewGraphqlCreateOK() *GraphqlCreateOK {
	return &GraphqlCreateOK{}
}

/* GraphqlCreateOK describes a response with status code 200, with default header values.

GraphqlCreateOK graphql create o k
*/
type GraphqlCreateOK struct {
	Payload *GraphqlCreateOKBody
}

func (o *GraphqlCreateOK) Error() string {
	return fmt.Sprintf("[POST /graphql/][%d] graphqlCreateOK  %+v", 200, o.Payload)
}
func (o *GraphqlCreateOK) GetPayload() *GraphqlCreateOKBody {
	return o.Payload
}

func (o *GraphqlCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GraphqlCreateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGraphqlCreateBadRequest creates a GraphqlCreateBadRequest with default headers values
func NewGraphqlCreateBadRequest() *GraphqlCreateBadRequest {
	return &GraphqlCreateBadRequest{}
}

/* GraphqlCreateBadRequest describes a response with status code 400, with default header values.

GraphqlCreateBadRequest graphql create bad request
*/
type GraphqlCreateBadRequest struct {
	Payload *GraphqlCreateBadRequestBody
}

func (o *GraphqlCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /graphql/][%d] graphqlCreateBadRequest  %+v", 400, o.Payload)
}
func (o *GraphqlCreateBadRequest) GetPayload() *GraphqlCreateBadRequestBody {
	return o.Payload
}

func (o *GraphqlCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GraphqlCreateBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GraphqlCreateBadRequestBody graphql create bad request body
swagger:model GraphqlCreateBadRequestBody
*/
type GraphqlCreateBadRequestBody struct {

	// errors
	Errors []interface{} `json:"errors"`
}

// Validate validates this graphql create bad request body
func (o *GraphqlCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this graphql create bad request body based on context it is used
func (o *GraphqlCreateBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GraphqlCreateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GraphqlCreateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GraphqlCreateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GraphqlCreateOKBody graphql create o k body
swagger:model GraphqlCreateOKBody
*/
type GraphqlCreateOKBody struct {

	// data
	Data interface{} `json:"data,omitempty"`
}

// Validate validates this graphql create o k body
func (o *GraphqlCreateOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this graphql create o k body based on context it is used
func (o *GraphqlCreateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GraphqlCreateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GraphqlCreateOKBody) UnmarshalBinary(b []byte) error {
	var res GraphqlCreateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
