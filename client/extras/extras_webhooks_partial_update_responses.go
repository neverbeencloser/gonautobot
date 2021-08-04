package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasWebhooksPartialUpdateReader is a Reader for the ExtrasWebhooksPartialUpdate structure.
type ExtrasWebhooksPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasWebhooksPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasWebhooksPartialUpdateOK creates a ExtrasWebhooksPartialUpdateOK with default headers values
func NewExtrasWebhooksPartialUpdateOK() *ExtrasWebhooksPartialUpdateOK {
	return &ExtrasWebhooksPartialUpdateOK{}
}

/* ExtrasWebhooksPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasWebhooksPartialUpdateOK extras webhooks partial update o k
*/
type ExtrasWebhooksPartialUpdateOK struct {
	Payload *models.Webhook
}

func (o *ExtrasWebhooksPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/webhooks/{id}/][%d] extrasWebhooksPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasWebhooksPartialUpdateOK) GetPayload() *models.Webhook {
	return o.Payload
}

func (o *ExtrasWebhooksPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Webhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
