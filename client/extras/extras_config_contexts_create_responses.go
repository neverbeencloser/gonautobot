package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasConfigContextsCreateReader is a Reader for the ExtrasConfigContextsCreate structure.
type ExtrasConfigContextsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasConfigContextsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextsCreateCreated creates a ExtrasConfigContextsCreateCreated with default headers values
func NewExtrasConfigContextsCreateCreated() *ExtrasConfigContextsCreateCreated {
	return &ExtrasConfigContextsCreateCreated{}
}

/* ExtrasConfigContextsCreateCreated describes a response with status code 201, with default header values.

ExtrasConfigContextsCreateCreated extras config contexts create created
*/
type ExtrasConfigContextsCreateCreated struct {
	Payload *models.ConfigContext
}

func (o *ExtrasConfigContextsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/config-contexts/][%d] extrasConfigContextsCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasConfigContextsCreateCreated) GetPayload() *models.ConfigContext {
	return o.Payload
}

func (o *ExtrasConfigContextsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
