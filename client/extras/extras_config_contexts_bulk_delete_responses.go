package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasConfigContextsBulkDeleteReader is a Reader for the ExtrasConfigContextsBulkDelete structure.
type ExtrasConfigContextsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasConfigContextsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextsBulkDeleteNoContent creates a ExtrasConfigContextsBulkDeleteNoContent with default headers values
func NewExtrasConfigContextsBulkDeleteNoContent() *ExtrasConfigContextsBulkDeleteNoContent {
	return &ExtrasConfigContextsBulkDeleteNoContent{}
}

/* ExtrasConfigContextsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasConfigContextsBulkDeleteNoContent extras config contexts bulk delete no content
*/
type ExtrasConfigContextsBulkDeleteNoContent struct {
}

func (o *ExtrasConfigContextsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/config-contexts/][%d] extrasConfigContextsBulkDeleteNoContent ", 204)
}

func (o *ExtrasConfigContextsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
