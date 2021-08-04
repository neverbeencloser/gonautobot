package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasComputedFieldsBulkPartialUpdateReader is a Reader for the ExtrasComputedFieldsBulkPartialUpdate structure.
type ExtrasComputedFieldsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasComputedFieldsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasComputedFieldsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasComputedFieldsBulkPartialUpdateOK creates a ExtrasComputedFieldsBulkPartialUpdateOK with default headers values
func NewExtrasComputedFieldsBulkPartialUpdateOK() *ExtrasComputedFieldsBulkPartialUpdateOK {
	return &ExtrasComputedFieldsBulkPartialUpdateOK{}
}

/* ExtrasComputedFieldsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasComputedFieldsBulkPartialUpdateOK extras computed fields bulk partial update o k
*/
type ExtrasComputedFieldsBulkPartialUpdateOK struct {
	Payload *models.ComputedField
}

func (o *ExtrasComputedFieldsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/computed-fields/][%d] extrasComputedFieldsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasComputedFieldsBulkPartialUpdateOK) GetPayload() *models.ComputedField {
	return o.Payload
}

func (o *ExtrasComputedFieldsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputedField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
