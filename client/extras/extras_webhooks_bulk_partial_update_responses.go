package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasWebhooksBulkPartialUpdateReader is a Reader for the ExtrasWebhooksBulkPartialUpdate structure.
type ExtrasWebhooksBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasWebhooksBulkPartialUpdateOK creates a ExtrasWebhooksBulkPartialUpdateOK with default headers values
func NewExtrasWebhooksBulkPartialUpdateOK() *ExtrasWebhooksBulkPartialUpdateOK {
	return &ExtrasWebhooksBulkPartialUpdateOK{}
}

/* ExtrasWebhooksBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksBulkPartialUpdateOK extras webhooks bulk partial update o k
*/
type ExtrasWebhooksBulkPartialUpdateOK struct {
	Payload *models.Webhook
}

func (o *ExtrasWebhooksBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/webhooks/][%d] extrasWebhooksBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasWebhooksBulkPartialUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
