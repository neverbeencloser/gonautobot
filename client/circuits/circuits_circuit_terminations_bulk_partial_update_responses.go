package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitTerminationsBulkPartialUpdateReader is a Reader for the CircuitsCircuitTerminationsBulkPartialUpdate structure.
type CircuitsCircuitTerminationsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsBulkPartialUpdateOK creates a CircuitsCircuitTerminationsBulkPartialUpdateOK with default headers values
func NewCircuitsCircuitTerminationsBulkPartialUpdateOK() *CircuitsCircuitTerminationsBulkPartialUpdateOK {
	return &CircuitsCircuitTerminationsBulkPartialUpdateOK{}
}

/* CircuitsCircuitTerminationsBulkPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsBulkPartialUpdateOK circuits circuit terminations bulk partial update o k
*/
type CircuitsCircuitTerminationsBulkPartialUpdateOK struct {
	Payload *models.CircuitTermination
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
