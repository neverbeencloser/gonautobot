package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasImageAttachmentsBulkDeleteReader is a Reader for the ExtrasImageAttachmentsBulkDelete structure.
type ExtrasImageAttachmentsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasImageAttachmentsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasImageAttachmentsBulkDeleteNoContent creates a ExtrasImageAttachmentsBulkDeleteNoContent with default headers values
func NewExtrasImageAttachmentsBulkDeleteNoContent() *ExtrasImageAttachmentsBulkDeleteNoContent {
	return &ExtrasImageAttachmentsBulkDeleteNoContent{}
}

/* ExtrasImageAttachmentsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasImageAttachmentsBulkDeleteNoContent extras image attachments bulk delete no content
*/
type ExtrasImageAttachmentsBulkDeleteNoContent struct {
}

func (o *ExtrasImageAttachmentsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/][%d] extrasImageAttachmentsBulkDeleteNoContent ", 204)
}

func (o *ExtrasImageAttachmentsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
