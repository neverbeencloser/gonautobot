package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldChoicesPartialUpdateReader is a Reader for the ExtrasCustomFieldChoicesPartialUpdate structure.
type ExtrasCustomFieldChoicesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoicesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldChoicesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldChoicesPartialUpdateOK creates a ExtrasCustomFieldChoicesPartialUpdateOK with default headers values
func NewExtrasCustomFieldChoicesPartialUpdateOK() *ExtrasCustomFieldChoicesPartialUpdateOK {
	return &ExtrasCustomFieldChoicesPartialUpdateOK{}
}

/* ExtrasCustomFieldChoicesPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldChoicesPartialUpdateOK extras custom field choices partial update o k
*/
type ExtrasCustomFieldChoicesPartialUpdateOK struct {
	Payload *models.CustomFieldChoice
}

func (o *ExtrasCustomFieldChoicesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-field-choices/{id}/][%d] extrasCustomFieldChoicesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomFieldChoicesPartialUpdateOK) GetPayload() *models.CustomFieldChoice {
	return o.Payload
}

func (o *ExtrasCustomFieldChoicesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomFieldChoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
