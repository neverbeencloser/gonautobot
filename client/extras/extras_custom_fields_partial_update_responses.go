package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldsPartialUpdateReader is a Reader for the ExtrasCustomFieldsPartialUpdate structure.
type ExtrasCustomFieldsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldsPartialUpdateOK creates a ExtrasCustomFieldsPartialUpdateOK with default headers values
func NewExtrasCustomFieldsPartialUpdateOK() *ExtrasCustomFieldsPartialUpdateOK {
	return &ExtrasCustomFieldsPartialUpdateOK{}
}

/* ExtrasCustomFieldsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldsPartialUpdateOK extras custom fields partial update o k
*/
type ExtrasCustomFieldsPartialUpdateOK struct {
	Payload *models.CustomField
}

func (o *ExtrasCustomFieldsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/custom-fields/{id}/][%d] extrasCustomFieldsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasCustomFieldsPartialUpdateOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
