package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldsCreateReader is a Reader for the ExtrasCustomFieldsCreate structure.
type ExtrasCustomFieldsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasCustomFieldsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldsCreateCreated creates a ExtrasCustomFieldsCreateCreated with default headers values
func NewExtrasCustomFieldsCreateCreated() *ExtrasCustomFieldsCreateCreated {
	return &ExtrasCustomFieldsCreateCreated{}
}

/* ExtrasCustomFieldsCreateCreated describes a response with status code 201, with default header values.

ExtrasCustomFieldsCreateCreated extras custom fields create created
*/
type ExtrasCustomFieldsCreateCreated struct {
	Payload *models.CustomField
}

func (o *ExtrasCustomFieldsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/custom-fields/][%d] extrasCustomFieldsCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasCustomFieldsCreateCreated) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
