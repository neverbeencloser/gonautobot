package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasImageAttachmentsDeleteReader is a Reader for the ExtrasImageAttachmentsDelete structure.
type ExtrasImageAttachmentsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasImageAttachmentsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasImageAttachmentsDeleteNoContent creates a ExtrasImageAttachmentsDeleteNoContent with default headers values
func NewExtrasImageAttachmentsDeleteNoContent() *ExtrasImageAttachmentsDeleteNoContent {
	return &ExtrasImageAttachmentsDeleteNoContent{}
}

/* ExtrasImageAttachmentsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasImageAttachmentsDeleteNoContent extras image attachments delete no content
*/
type ExtrasImageAttachmentsDeleteNoContent struct {
}

func (o *ExtrasImageAttachmentsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extrasImageAttachmentsDeleteNoContent ", 204)
}

func (o *ExtrasImageAttachmentsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
