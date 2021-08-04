package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasJobResultsBulkDeleteReader is a Reader for the ExtrasJobResultsBulkDelete structure.
type ExtrasJobResultsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJobResultsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasJobResultsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasJobResultsBulkDeleteNoContent creates a ExtrasJobResultsBulkDeleteNoContent with default headers values
func NewExtrasJobResultsBulkDeleteNoContent() *ExtrasJobResultsBulkDeleteNoContent {
	return &ExtrasJobResultsBulkDeleteNoContent{}
}

/* ExtrasJobResultsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasJobResultsBulkDeleteNoContent extras job results bulk delete no content
*/
type ExtrasJobResultsBulkDeleteNoContent struct {
}

func (o *ExtrasJobResultsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/job-results/][%d] extrasJobResultsBulkDeleteNoContent ", 204)
}

func (o *ExtrasJobResultsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
