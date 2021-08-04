package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldsBulkPartialUpdateReader is a Reader for the ExtrasCustomFieldsBulkPartialUpdate structure.
type ExtrasCustomFieldsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldsBulkPartialUpdateOK creates a ExtrasCustomFieldsBulkPartialUpdateOK with default headers values
func NewExtrasCustomFieldsBulkPartialUpdateOK() *ExtrasCustomFieldsBulkPartialUpdateOK {
	return &ExtrasCustomFieldsBulkPartialUpdateOK{}
}

/* ExtrasCustomFieldsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldsBulkPartialUpdateOK extras custom fields bulk partial update o k
*/
type ExtrasCustomFieldsBulkPartialUpdateOK struct {
	Payload *models.CustomField
}

func (o *ExtrasCustomFieldsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-fields/][%d] extrasCustomFieldsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomFieldsBulkPartialUpdateOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
