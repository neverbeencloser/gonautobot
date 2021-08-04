package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasWebhooksDeleteReader is a Reader for the ExtrasWebhooksDelete structure.
type ExtrasWebhooksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasWebhooksDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasWebhooksDeleteNoContent creates a ExtrasWebhooksDeleteNoContent with default headers values
func NewExtrasWebhooksDeleteNoContent() *ExtrasWebhooksDeleteNoContent {
	return &ExtrasWebhooksDeleteNoContent{}
}

/* ExtrasWebhooksDeleteNoContent describes a response with status code 204, with default header values.

ExtrasWebhooksDeleteNoContent extras webhooks delete no content
*/
type ExtrasWebhooksDeleteNoContent struct {
}

func (o *ExtrasWebhooksDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/webhooks/{id}/][%d] extrasWebhooksDeleteNoContent ", 204)
}

func (o *ExtrasWebhooksDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
