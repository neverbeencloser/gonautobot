package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasStatusesBulkDeleteReader is a Reader for the ExtrasStatusesBulkDelete structure.
type ExtrasStatusesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasStatusesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasStatusesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasStatusesBulkDeleteNoContent creates a ExtrasStatusesBulkDeleteNoContent with default headers values
func NewExtrasStatusesBulkDeleteNoContent() *ExtrasStatusesBulkDeleteNoContent {
	return &ExtrasStatusesBulkDeleteNoContent{}
}

/* ExtrasStatusesBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasStatusesBulkDeleteNoContent extras statuses bulk delete no content
*/
type ExtrasStatusesBulkDeleteNoContent struct {
}

func (o *ExtrasStatusesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/statuses/][%d] extrasStatusesBulkDeleteNoContent ", 204)
}

func (o *ExtrasStatusesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
