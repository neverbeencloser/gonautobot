package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasConfigContextSchemasBulkPartialUpdateReader is a Reader for the ExtrasConfigContextSchemasBulkPartialUpdate structure.
type ExtrasConfigContextSchemasBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextSchemasBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextSchemasBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextSchemasBulkPartialUpdateOK creates a ExtrasConfigContextSchemasBulkPartialUpdateOK with default headers values
func NewExtrasConfigContextSchemasBulkPartialUpdateOK() *ExtrasConfigContextSchemasBulkPartialUpdateOK {
	return &ExtrasConfigContextSchemasBulkPartialUpdateOK{}
}

/* ExtrasConfigContextSchemasBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextSchemasBulkPartialUpdateOK extras config context schemas bulk partial update o k
*/
type ExtrasConfigContextSchemasBulkPartialUpdateOK struct {
	Payload *models.ConfigContextSchema
}

func (o *ExtrasConfigContextSchemasBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/config-context-schemas/][%d] extrasConfigContextSchemasBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasConfigContextSchemasBulkPartialUpdateOK) GetPayload() *models.ConfigContextSchema {
	return o.Payload
}

func (o *ExtrasConfigContextSchemasBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContextSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
