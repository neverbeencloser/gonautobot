package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasWebhooksBulkDeleteReader is a Reader for the ExtrasWebhooksBulkDelete structure.
type ExtrasWebhooksBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasWebhooksBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasWebhooksBulkDeleteNoContent creates a ExtrasWebhooksBulkDeleteNoContent with default headers values
func NewExtrasWebhooksBulkDeleteNoContent() *ExtrasWebhooksBulkDeleteNoContent {
	return &ExtrasWebhooksBulkDeleteNoContent{}
}

/* ExtrasWebhooksBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasWebhooksBulkDeleteNoContent extras webhooks bulk delete no content
*/
type ExtrasWebhooksBulkDeleteNoContent struct {
}

func (o *ExtrasWebhooksBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/webhooks/][%d] extrasWebhooksBulkDeleteNoContent ", 204)
}

func (o *ExtrasWebhooksBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
