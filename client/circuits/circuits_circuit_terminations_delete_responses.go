package circuits

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CircuitsCircuitTerminationsDeleteReader is a Reader for the CircuitsCircuitTerminationsDelete structure.
type CircuitsCircuitTerminationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsCircuitTerminationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsDeleteNoContent creates a CircuitsCircuitTerminationsDeleteNoContent with default headers values
func NewCircuitsCircuitTerminationsDeleteNoContent() *CircuitsCircuitTerminationsDeleteNoContent {
	return &CircuitsCircuitTerminationsDeleteNoContent{}
}

/* CircuitsCircuitTerminationsDeleteNoContent describes a response with status code 204, with default header values.

CircuitsCircuitTerminationsDeleteNoContent circuits circuit terminations delete no content
*/
type CircuitsCircuitTerminationsDeleteNoContent struct {
}

func (o *CircuitsCircuitTerminationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsDeleteNoContent ", 204)
}

func (o *CircuitsCircuitTerminationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
