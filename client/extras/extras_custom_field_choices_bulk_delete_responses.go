package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasCustomFieldChoicesBulkDeleteReader is a Reader for the ExtrasCustomFieldChoicesBulkDelete structure.
type ExtrasCustomFieldChoicesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoicesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasCustomFieldChoicesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldChoicesBulkDeleteNoContent creates a ExtrasCustomFieldChoicesBulkDeleteNoContent with default headers values
func NewExtrasCustomFieldChoicesBulkDeleteNoContent() *ExtrasCustomFieldChoicesBulkDeleteNoContent {
	return &ExtrasCustomFieldChoicesBulkDeleteNoContent{}
}

/* ExtrasCustomFieldChoicesBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasCustomFieldChoicesBulkDeleteNoContent extras custom field choices bulk delete no content
*/
type ExtrasCustomFieldChoicesBulkDeleteNoContent struct {
}

func (o *ExtrasCustomFieldChoicesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-field-choices/][%d] extrasCustomFieldChoicesBulkDeleteNoContent ", 204)
}

func (o *ExtrasCustomFieldChoicesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
