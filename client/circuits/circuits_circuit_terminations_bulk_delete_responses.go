package circuits

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CircuitsCircuitTerminationsBulkDeleteReader is a Reader for the CircuitsCircuitTerminationsBulkDelete structure.
type CircuitsCircuitTerminationsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsCircuitTerminationsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsBulkDeleteNoContent creates a CircuitsCircuitTerminationsBulkDeleteNoContent with default headers values
func NewCircuitsCircuitTerminationsBulkDeleteNoContent() *CircuitsCircuitTerminationsBulkDeleteNoContent {
	return &CircuitsCircuitTerminationsBulkDeleteNoContent{}
}

/* CircuitsCircuitTerminationsBulkDeleteNoContent describes a response with status code 204, with default header values.

CircuitsCircuitTerminationsBulkDeleteNoContent circuits circuit terminations bulk delete no content
*/
type CircuitsCircuitTerminationsBulkDeleteNoContent struct {
}

func (o *CircuitsCircuitTerminationsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsBulkDeleteNoContent ", 204)
}

func (o *CircuitsCircuitTerminationsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
