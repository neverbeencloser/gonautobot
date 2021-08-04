package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitTypesReadReader is a Reader for the CircuitsCircuitTypesRead structure.
type CircuitsCircuitTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTypesReadOK creates a CircuitsCircuitTypesReadOK with default headers values
func NewCircuitsCircuitTypesReadOK() *CircuitsCircuitTypesReadOK {
	return &CircuitsCircuitTypesReadOK{}
}

/* CircuitsCircuitTypesReadOK describes a response with status code 200, with default header values.

CircuitsCircuitTypesReadOK circuits circuit types read o k
*/
type CircuitsCircuitTypesReadOK struct {
	Payload *models.CircuitType
}

func (o *CircuitsCircuitTypesReadOK) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesReadOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTypesReadOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
