package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasCustomFieldChoicesDeleteReader is a Reader for the ExtrasCustomFieldChoicesDelete structure.
type ExtrasCustomFieldChoicesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldChoicesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasCustomFieldChoicesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldChoicesDeleteNoContent creates a ExtrasCustomFieldChoicesDeleteNoContent with default headers values
func NewExtrasCustomFieldChoicesDeleteNoContent() *ExtrasCustomFieldChoicesDeleteNoContent {
	return &ExtrasCustomFieldChoicesDeleteNoContent{}
}

/* ExtrasCustomFieldChoicesDeleteNoContent describes a response with status code 204, with default header values.

ExtrasCustomFieldChoicesDeleteNoContent extras custom field choices delete no content
*/
type ExtrasCustomFieldChoicesDeleteNoContent struct {
}

func (o *ExtrasCustomFieldChoicesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-field-choices/{id}/][%d] extrasCustomFieldChoicesDeleteNoContent ", 204)
}

func (o *ExtrasCustomFieldChoicesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
