package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldsReadReader is a Reader for the ExtrasCustomFieldsRead structure.
type ExtrasCustomFieldsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldsReadOK creates a ExtrasCustomFieldsReadOK with default headers values
func NewExtrasCustomFieldsReadOK() *ExtrasCustomFieldsReadOK {
	return &ExtrasCustomFieldsReadOK{}
}

/* ExtrasCustomFieldsReadOK describes a response with status code 200, with default header values.

ExtrasCustomFieldsReadOK extras custom fields read o k
*/
type ExtrasCustomFieldsReadOK struct {
	Payload *models.CustomField
}

func (o *ExtrasCustomFieldsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/custom-fields/{id}/][%d] extrasCustomFieldsReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomFieldsReadOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
