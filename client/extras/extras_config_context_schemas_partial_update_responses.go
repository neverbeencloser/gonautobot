package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasConfigContextSchemasPartialUpdateReader is a Reader for the ExtrasConfigContextSchemasPartialUpdate structure.
type ExtrasConfigContextSchemasPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextSchemasPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextSchemasPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextSchemasPartialUpdateOK creates a ExtrasConfigContextSchemasPartialUpdateOK with default headers values
func NewExtrasConfigContextSchemasPartialUpdateOK() *ExtrasConfigContextSchemasPartialUpdateOK {
	return &ExtrasConfigContextSchemasPartialUpdateOK{}
}

/* ExtrasConfigContextSchemasPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextSchemasPartialUpdateOK extras config context schemas partial update o k
*/
type ExtrasConfigContextSchemasPartialUpdateOK struct {
	Payload *models.ConfigContextSchema
}

func (o *ExtrasConfigContextSchemasPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-context-schemas/{id}/][%d] extrasConfigContextSchemasPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasConfigContextSchemasPartialUpdateOK) GetPayload() *models.ConfigContextSchema {
	return o.Payload
}

func (o *ExtrasConfigContextSchemasPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContextSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
