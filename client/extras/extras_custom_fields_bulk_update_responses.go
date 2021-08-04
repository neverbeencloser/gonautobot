package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldsBulkUpdateReader is a Reader for the ExtrasCustomFieldsBulkUpdate structure.
type ExtrasCustomFieldsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldsBulkUpdateOK creates a ExtrasCustomFieldsBulkUpdateOK with default headers values
func NewExtrasCustomFieldsBulkUpdateOK() *ExtrasCustomFieldsBulkUpdateOK {
	return &ExtrasCustomFieldsBulkUpdateOK{}
}

/* ExtrasCustomFieldsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldsBulkUpdateOK extras custom fields bulk update o k
*/
type ExtrasCustomFieldsBulkUpdateOK struct {
	Payload *models.CustomField
}

func (o *ExtrasCustomFieldsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-fields/][%d] extrasCustomFieldsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomFieldsBulkUpdateOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
