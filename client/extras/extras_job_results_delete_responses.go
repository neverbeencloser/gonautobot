package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasJobResultsDeleteReader is a Reader for the ExtrasJobResultsDelete structure.
type ExtrasJobResultsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJobResultsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasJobResultsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasJobResultsDeleteNoContent creates a ExtrasJobResultsDeleteNoContent with default headers values
func NewExtrasJobResultsDeleteNoContent() *ExtrasJobResultsDeleteNoContent {
	return &ExtrasJobResultsDeleteNoContent{}
}

/* ExtrasJobResultsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasJobResultsDeleteNoContent extras job results delete no content
*/
type ExtrasJobResultsDeleteNoContent struct {
}

func (o *ExtrasJobResultsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/job-results/{id}/][%d] extrasJobResultsDeleteNoContent ", 204)
}

func (o *ExtrasJobResultsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
