package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasConfigContextSchemasBulkUpdateReader is a Reader for the ExtrasConfigContextSchemasBulkUpdate structure.
type ExtrasConfigContextSchemasBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextSchemasBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasConfigContextSchemasBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextSchemasBulkUpdateOK creates a ExtrasConfigContextSchemasBulkUpdateOK with default headers values
func NewExtrasConfigContextSchemasBulkUpdateOK() *ExtrasConfigContextSchemasBulkUpdateOK {
	return &ExtrasConfigContextSchemasBulkUpdateOK{}
}

/* ExtrasConfigContextSchemasBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasConfigContextSchemasBulkUpdateOK extras config context schemas bulk update o k
*/
type ExtrasConfigContextSchemasBulkUpdateOK struct {
	Payload *models.ConfigContextSchema
}

func (o *ExtrasConfigContextSchemasBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/config-context-schemas/][%d] extrasConfigContextSchemasBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasConfigContextSchemasBulkUpdateOK) GetPayload() *models.ConfigContextSchema {
	return o.Payload
}

func (o *ExtrasConfigContextSchemasBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContextSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
