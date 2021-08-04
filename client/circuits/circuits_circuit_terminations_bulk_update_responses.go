package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitTerminationsBulkUpdateReader is a Reader for the CircuitsCircuitTerminationsBulkUpdate structure.
type CircuitsCircuitTerminationsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTerminationsBulkUpdateOK creates a CircuitsCircuitTerminationsBulkUpdateOK with default headers values
func NewCircuitsCircuitTerminationsBulkUpdateOK() *CircuitsCircuitTerminationsBulkUpdateOK {
	return &CircuitsCircuitTerminationsBulkUpdateOK{}
}

/* CircuitsCircuitTerminationsBulkUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsBulkUpdateOK circuits circuit terminations bulk update o k
*/
type CircuitsCircuitTerminationsBulkUpdateOK struct {
	Payload *models.CircuitTermination
}

func (o *CircuitsCircuitTerminationsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-terminations/][%d] circuitsCircuitTerminationsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTerminationsBulkUpdateOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
