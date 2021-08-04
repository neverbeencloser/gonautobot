package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasComputedFieldsBulkDeleteReader is a Reader for the ExtrasComputedFieldsBulkDelete structure.
type ExtrasComputedFieldsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasComputedFieldsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasComputedFieldsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasComputedFieldsBulkDeleteNoContent creates a ExtrasComputedFieldsBulkDeleteNoContent with default headers values
func NewExtrasComputedFieldsBulkDeleteNoContent() *ExtrasComputedFieldsBulkDeleteNoContent {
	return &ExtrasComputedFieldsBulkDeleteNoContent{}
}

/* ExtrasComputedFieldsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasComputedFieldsBulkDeleteNoContent extras computed fields bulk delete no content
*/
type ExtrasComputedFieldsBulkDeleteNoContent struct {
}

func (o *ExtrasComputedFieldsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/computed-fields/][%d] extrasComputedFieldsBulkDeleteNoContent ", 204)
}

func (o *ExtrasComputedFieldsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
