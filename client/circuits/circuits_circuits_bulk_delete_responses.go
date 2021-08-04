package circuits

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CircuitsCircuitsBulkDeleteReader is a Reader for the CircuitsCircuitsBulkDelete structure.
type CircuitsCircuitsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsCircuitsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitsBulkDeleteNoContent creates a CircuitsCircuitsBulkDeleteNoContent with default headers values
func NewCircuitsCircuitsBulkDeleteNoContent() *CircuitsCircuitsBulkDeleteNoContent {
	return &CircuitsCircuitsBulkDeleteNoContent{}
}

/* CircuitsCircuitsBulkDeleteNoContent describes a response with status code 204, with default header values.

CircuitsCircuitsBulkDeleteNoContent circuits circuits bulk delete no content
*/
type CircuitsCircuitsBulkDeleteNoContent struct {
}

func (o *CircuitsCircuitsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/circuits/][%d] circuitsCircuitsBulkDeleteNoContent ", 204)
}

func (o *CircuitsCircuitsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
