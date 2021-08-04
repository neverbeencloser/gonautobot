package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitTypesPartialUpdateReader is a Reader for the CircuitsCircuitTypesPartialUpdate structure.
type CircuitsCircuitTypesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTypesPartialUpdateOK creates a CircuitsCircuitTypesPartialUpdateOK with default headers values
func NewCircuitsCircuitTypesPartialUpdateOK() *CircuitsCircuitTypesPartialUpdateOK {
	return &CircuitsCircuitTypesPartialUpdateOK{}
}

/* CircuitsCircuitTypesPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTypesPartialUpdateOK circuits circuit types partial update o k
*/
type CircuitsCircuitTypesPartialUpdateOK struct {
	Payload *models.CircuitType
}

func (o *CircuitsCircuitTypesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTypesPartialUpdateOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
