package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasConfigContextSchemasBulkDeleteReader is a Reader for the ExtrasConfigContextSchemasBulkDelete structure.
type ExtrasConfigContextSchemasBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextSchemasBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasConfigContextSchemasBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextSchemasBulkDeleteNoContent creates a ExtrasConfigContextSchemasBulkDeleteNoContent with default headers values
func NewExtrasConfigContextSchemasBulkDeleteNoContent() *ExtrasConfigContextSchemasBulkDeleteNoContent {
	return &ExtrasConfigContextSchemasBulkDeleteNoContent{}
}

/* ExtrasConfigContextSchemasBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasConfigContextSchemasBulkDeleteNoContent extras config context schemas bulk delete no content
*/
type ExtrasConfigContextSchemasBulkDeleteNoContent struct {
}

func (o *ExtrasConfigContextSchemasBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/config-context-schemas/][%d] extrasConfigContextSchemasBulkDeleteNoContent ", 204)
}

func (o *ExtrasConfigContextSchemasBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
