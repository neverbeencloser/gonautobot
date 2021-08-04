package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasComputedFieldsBulkUpdateReader is a Reader for the ExtrasComputedFieldsBulkUpdate structure.
type ExtrasComputedFieldsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasComputedFieldsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasComputedFieldsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasComputedFieldsBulkUpdateOK creates a ExtrasComputedFieldsBulkUpdateOK with default headers values
func NewExtrasComputedFieldsBulkUpdateOK() *ExtrasComputedFieldsBulkUpdateOK {
	return &ExtrasComputedFieldsBulkUpdateOK{}
}

/* ExtrasComputedFieldsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasComputedFieldsBulkUpdateOK extras computed fields bulk update o k
*/
type ExtrasComputedFieldsBulkUpdateOK struct {
	Payload *models.ComputedField
}

func (o *ExtrasComputedFieldsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/computed-fields/][%d] extrasComputedFieldsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasComputedFieldsBulkUpdateOK) GetPayload() *models.ComputedField {
	return o.Payload
}

func (o *ExtrasComputedFieldsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputedField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
