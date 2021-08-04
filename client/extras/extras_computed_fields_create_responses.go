package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasComputedFieldsCreateReader is a Reader for the ExtrasComputedFieldsCreate structure.
type ExtrasComputedFieldsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasComputedFieldsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasComputedFieldsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasComputedFieldsCreateCreated creates a ExtrasComputedFieldsCreateCreated with default headers values
func NewExtrasComputedFieldsCreateCreated() *ExtrasComputedFieldsCreateCreated {
	return &ExtrasComputedFieldsCreateCreated{}
}

/* ExtrasComputedFieldsCreateCreated describes a response with status code 201, with default header values.

ExtrasComputedFieldsCreateCreated extras computed fields create created
*/
type ExtrasComputedFieldsCreateCreated struct {
	Payload *models.ComputedField
}

func (o *ExtrasComputedFieldsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/computed-fields/][%d] extrasComputedFieldsCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasComputedFieldsCreateCreated) GetPayload() *models.ComputedField {
	return o.Payload
}

func (o *ExtrasComputedFieldsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputedField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
