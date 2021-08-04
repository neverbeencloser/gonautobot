package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasTagsBulkDeleteReader is a Reader for the ExtrasTagsBulkDelete structure.
type ExtrasTagsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasTagsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasTagsBulkDeleteNoContent creates a ExtrasTagsBulkDeleteNoContent with default headers values
func NewExtrasTagsBulkDeleteNoContent() *ExtrasTagsBulkDeleteNoContent {
	return &ExtrasTagsBulkDeleteNoContent{}
}

/* ExtrasTagsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasTagsBulkDeleteNoContent extras tags bulk delete no content
*/
type ExtrasTagsBulkDeleteNoContent struct {
}

func (o *ExtrasTagsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/tags/][%d] extrasTagsBulkDeleteNoContent ", 204)
}

func (o *ExtrasTagsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
