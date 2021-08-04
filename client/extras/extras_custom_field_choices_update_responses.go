package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldChoicesUpdateReader is a Reader for the ExtrasCustomFieldChoicesUpdate structure.
type ExtrasCustomFieldChoicesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoicesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldChoicesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldChoicesUpdateOK creates a ExtrasCustomFieldChoicesUpdateOK with default headers values
func NewExtrasCustomFieldChoicesUpdateOK() *ExtrasCustomFieldChoicesUpdateOK {
	return &ExtrasCustomFieldChoicesUpdateOK{}
}

/* ExtrasCustomFieldChoicesUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldChoicesUpdateOK extras custom field choices update o k
*/
type ExtrasCustomFieldChoicesUpdateOK struct {
	Payload *models.CustomFieldChoice
}

func (o *ExtrasCustomFieldChoicesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-field-choices/{id}/][%d] extrasCustomFieldChoicesUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomFieldChoicesUpdateOK) GetPayload() *models.CustomFieldChoice {
	return o.Payload
}

func (o *ExtrasCustomFieldChoicesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomFieldChoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
