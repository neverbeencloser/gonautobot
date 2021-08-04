package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasComputedFieldsDeleteReader is a Reader for the ExtrasComputedFieldsDelete structure.
type ExtrasComputedFieldsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasComputedFieldsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasComputedFieldsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasComputedFieldsDeleteNoContent creates a ExtrasComputedFieldsDeleteNoContent with default headers values
func NewExtrasComputedFieldsDeleteNoContent() *ExtrasComputedFieldsDeleteNoContent {
	return &ExtrasComputedFieldsDeleteNoContent{}
}

/* ExtrasComputedFieldsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasComputedFieldsDeleteNoContent extras computed fields delete no content
*/
type ExtrasComputedFieldsDeleteNoContent struct {
}

func (o *ExtrasComputedFieldsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/computed-fields/{id}/][%d] extrasComputedFieldsDeleteNoContent ", 204)
}

func (o *ExtrasComputedFieldsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
