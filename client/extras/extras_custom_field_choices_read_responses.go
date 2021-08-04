package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldChoicesReadReader is a Reader for the ExtrasCustomFieldChoicesRead structure.
type ExtrasCustomFieldChoicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldChoicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldChoicesReadOK creates a ExtrasCustomFieldChoicesReadOK with default headers values
func NewExtrasCustomFieldChoicesReadOK() *ExtrasCustomFieldChoicesReadOK {
	return &ExtrasCustomFieldChoicesReadOK{}
}

/* ExtrasCustomFieldChoicesReadOK describes a response with status code 200, with default header values.

ExtrasCustomFieldChoicesReadOK extras custom field choices read o k
*/
type ExtrasCustomFieldChoicesReadOK struct {
	Payload *models.CustomFieldChoice
}

func (o *ExtrasCustomFieldChoicesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/custom-field-choices/{id}/][%d] extrasCustomFieldChoicesReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomFieldChoicesReadOK) GetPayload() *models.CustomFieldChoice {
	return o.Payload
}

func (o *ExtrasCustomFieldChoicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomFieldChoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
