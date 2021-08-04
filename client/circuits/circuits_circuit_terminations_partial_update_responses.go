package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitTerminationsPartialUpdateReader is a Reader for the CircuitsCircuitTerminationsPartialUpdate structure.
type CircuitsCircuitTerminationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsPartialUpdateOK creates a CircuitsCircuitTerminationsPartialUpdateOK with default headers values
func NewCircuitsCircuitTerminationsPartialUpdateOK() *CircuitsCircuitTerminationsPartialUpdateOK {
	return &CircuitsCircuitTerminationsPartialUpdateOK{}
}

/* CircuitsCircuitTerminationsPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsPartialUpdateOK circuits circuit terminations partial update o k
*/
type CircuitsCircuitTerminationsPartialUpdateOK struct {
	Payload *models.CircuitTermination
}

func (o *CircuitsCircuitTerminationsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTerminationsPartialUpdateOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
