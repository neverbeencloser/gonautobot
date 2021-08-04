package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasCustomFieldsBulkDeleteReader is a Reader for the ExtrasCustomFieldsBulkDelete structure.
type ExtrasCustomFieldsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasCustomFieldsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldsBulkDeleteNoContent creates a ExtrasCustomFieldsBulkDeleteNoContent with default headers values
func NewExtrasCustomFieldsBulkDeleteNoContent() *ExtrasCustomFieldsBulkDeleteNoContent {
	return &ExtrasCustomFieldsBulkDeleteNoContent{}
}

/* ExtrasCustomFieldsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasCustomFieldsBulkDeleteNoContent extras custom fields bulk delete no content
*/
type ExtrasCustomFieldsBulkDeleteNoContent struct {
}

func (o *ExtrasCustomFieldsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-fields/][%d] extrasCustomFieldsBulkDeleteNoContent ", 204)
}

func (o *ExtrasCustomFieldsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
