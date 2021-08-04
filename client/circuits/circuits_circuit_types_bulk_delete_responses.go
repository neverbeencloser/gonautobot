package circuits

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CircuitsCircuitTypesBulkDeleteReader is a Reader for the CircuitsCircuitTypesBulkDelete structure.
type CircuitsCircuitTypesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsCircuitTypesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTypesBulkDeleteNoContent creates a CircuitsCircuitTypesBulkDeleteNoContent with default headers values
func NewCircuitsCircuitTypesBulkDeleteNoContent() *CircuitsCircuitTypesBulkDeleteNoContent {
	return &CircuitsCircuitTypesBulkDeleteNoContent{}
}

/* CircuitsCircuitTypesBulkDeleteNoContent describes a response with status code 204, with default header values.

CircuitsCircuitTypesBulkDeleteNoContent circuits circuit types bulk delete no content
*/
type CircuitsCircuitTypesBulkDeleteNoContent struct {
}

func (o *CircuitsCircuitTypesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/circuit-types/][%d] circuitsCircuitTypesBulkDeleteNoContent ", 204)
}

func (o *CircuitsCircuitTypesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
