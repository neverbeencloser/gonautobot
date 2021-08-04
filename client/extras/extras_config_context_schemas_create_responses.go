package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasConfigContextSchemasCreateReader is a Reader for the ExtrasConfigContextSchemasCreate structure.
type ExtrasConfigContextSchemasCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextSchemasCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasConfigContextSchemasCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextSchemasCreateCreated creates a ExtrasConfigContextSchemasCreateCreated with default headers values
func NewExtrasConfigContextSchemasCreateCreated() *ExtrasConfigContextSchemasCreateCreated {
	return &ExtrasConfigContextSchemasCreateCreated{}
}

/* ExtrasConfigContextSchemasCreateCreated describes a response with status code 201, with default header values.

ExtrasConfigContextSchemasCreateCreated extras config context schemas create created
*/
type ExtrasConfigContextSchemasCreateCreated struct {
	Payload *models.ConfigContextSchema
}

func (o *ExtrasConfigContextSchemasCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/config-context-schemas/][%d] extrasConfigContextSchemasCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasConfigContextSchemasCreateCreated) GetPayload() *models.ConfigContextSchema {
	return o.Payload
}

func (o *ExtrasConfigContextSchemasCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContextSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
