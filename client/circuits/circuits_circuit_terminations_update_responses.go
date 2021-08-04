package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitTerminationsUpdateReader is a Reader for the CircuitsCircuitTerminationsUpdate structure.
type CircuitsCircuitTerminationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsUpdateOK creates a CircuitsCircuitTerminationsUpdateOK with default headers values
func NewCircuitsCircuitTerminationsUpdateOK() *CircuitsCircuitTerminationsUpdateOK {
	return &CircuitsCircuitTerminationsUpdateOK{}
}

/* CircuitsCircuitTerminationsUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsUpdateOK circuits circuit terminations update o k
*/
type CircuitsCircuitTerminationsUpdateOK struct {
	Payload *models.CircuitTermination
}

func (o *CircuitsCircuitTerminationsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTerminationsUpdateOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
