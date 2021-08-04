package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasConfigContextSchemasDeleteReader is a Reader for the ExtrasConfigContextSchemasDelete structure.
type ExtrasConfigContextSchemasDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextSchemasDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasConfigContextSchemasDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasConfigContextSchemasDeleteNoContent creates a ExtrasConfigContextSchemasDeleteNoContent with default headers values
func NewExtrasConfigContextSchemasDeleteNoContent() *ExtrasConfigContextSchemasDeleteNoContent {
	return &ExtrasConfigContextSchemasDeleteNoContent{}
}

/* ExtrasConfigContextSchemasDeleteNoContent describes a response with status code 204, with default header values.

ExtrasConfigContextSchemasDeleteNoContent extras config context schemas delete no content
*/
type ExtrasConfigContextSchemasDeleteNoContent struct {
}

func (o *ExtrasConfigContextSchemasDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/config-context-schemas/{id}/][%d] extrasConfigContextSchemasDeleteNoContent ", 204)
}

func (o *ExtrasConfigContextSchemasDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
