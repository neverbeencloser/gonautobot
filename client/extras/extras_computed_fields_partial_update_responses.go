package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasComputedFieldsPartialUpdateReader is a Reader for the ExtrasComputedFieldsPartialUpdate structure.
type ExtrasComputedFieldsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasComputedFieldsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasComputedFieldsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasComputedFieldsPartialUpdateOK creates a ExtrasComputedFieldsPartialUpdateOK with default headers values
func NewExtrasComputedFieldsPartialUpdateOK() *ExtrasComputedFieldsPartialUpdateOK {
	return &ExtrasComputedFieldsPartialUpdateOK{}
}

/* ExtrasComputedFieldsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasComputedFieldsPartialUpdateOK extras computed fields partial update o k
*/
type ExtrasComputedFieldsPartialUpdateOK struct {
	Payload *models.ComputedField
}

func (o *ExtrasComputedFieldsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/computed-fields/{id}/][%d] extrasComputedFieldsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasComputedFieldsPartialUpdateOK) GetPayload() *models.ComputedField {
	return o.Payload
}

func (o *ExtrasComputedFieldsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputedField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
