package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitTerminationsTraceReader is a Reader for the CircuitsCircuitTerminationsTrace structure.
type CircuitsCircuitTerminationsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsTraceOK creates a CircuitsCircuitTerminationsTraceOK with default headers values
func NewCircuitsCircuitTerminationsTraceOK() *CircuitsCircuitTerminationsTraceOK {
	return &CircuitsCircuitTerminationsTraceOK{}
}

/* CircuitsCircuitTerminationsTraceOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsTraceOK circuits circuit terminations trace o k
*/
type CircuitsCircuitTerminationsTraceOK struct {
	Payload *models.CircuitTermination
}

func (o *CircuitsCircuitTerminationsTraceOK) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/trace/][%d] circuitsCircuitTerminationsTraceOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTerminationsTraceOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
