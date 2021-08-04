package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasCustomFieldsDeleteReader is a Reader for the ExtrasCustomFieldsDelete structure.
type ExtrasCustomFieldsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasCustomFieldsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldsDeleteNoContent creates a ExtrasCustomFieldsDeleteNoContent with default headers values
func NewExtrasCustomFieldsDeleteNoContent() *ExtrasCustomFieldsDeleteNoContent {
	return &ExtrasCustomFieldsDeleteNoContent{}
}

/* ExtrasCustomFieldsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasCustomFieldsDeleteNoContent extras custom fields delete no content
*/
type ExtrasCustomFieldsDeleteNoContent struct {
}

func (o *ExtrasCustomFieldsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-fields/{id}/][%d] extrasCustomFieldsDeleteNoContent ", 204)
}

func (o *ExtrasCustomFieldsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
