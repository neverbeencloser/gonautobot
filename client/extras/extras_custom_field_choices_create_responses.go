package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasCustomFieldChoicesCreateReader is a Reader for the ExtrasCustomFieldChoicesCreate structure.
type ExtrasCustomFieldChoicesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoicesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasCustomFieldChoicesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldChoicesCreateCreated creates a ExtrasCustomFieldChoicesCreateCreated with default headers values
func NewExtrasCustomFieldChoicesCreateCreated() *ExtrasCustomFieldChoicesCreateCreated {
	return &ExtrasCustomFieldChoicesCreateCreated{}
}

/* ExtrasCustomFieldChoicesCreateCreated describes a response with status code 201, with default header values.

ExtrasCustomFieldChoicesCreateCreated extras custom field choices create created
*/
type ExtrasCustomFieldChoicesCreateCreated struct {
	Payload *models.CustomFieldChoice
}

func (o *ExtrasCustomFieldChoicesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/custom-field-choices/][%d] extrasCustomFieldChoicesCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasCustomFieldChoicesCreateCreated) GetPayload() *models.CustomFieldChoice {
	return o.Payload
}

func (o *ExtrasCustomFieldChoicesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomFieldChoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
