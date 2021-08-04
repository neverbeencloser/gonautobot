package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldChoicesBulkUpdateReader is a Reader for the ExtrasCustomFieldChoicesBulkUpdate structure.
type ExtrasCustomFieldChoicesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoicesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldChoicesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldChoicesBulkUpdateOK creates a ExtrasCustomFieldChoicesBulkUpdateOK with default headers values
func NewExtrasCustomFieldChoicesBulkUpdateOK() *ExtrasCustomFieldChoicesBulkUpdateOK {
	return &ExtrasCustomFieldChoicesBulkUpdateOK{}
}

/* ExtrasCustomFieldChoicesBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldChoicesBulkUpdateOK extras custom field choices bulk update o k
*/
type ExtrasCustomFieldChoicesBulkUpdateOK struct {
	Payload *models.CustomFieldChoice
}

func (o *ExtrasCustomFieldChoicesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-field-choices/][%d] extrasCustomFieldChoicesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomFieldChoicesBulkUpdateOK) GetPayload() *models.CustomFieldChoice {
	return o.Payload
}

func (o *ExtrasCustomFieldChoicesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomFieldChoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
