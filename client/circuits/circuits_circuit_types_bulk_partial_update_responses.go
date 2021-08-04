package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitTypesBulkPartialUpdateReader is a Reader for the CircuitsCircuitTypesBulkPartialUpdate structure.
type CircuitsCircuitTypesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTypesBulkPartialUpdateOK creates a CircuitsCircuitTypesBulkPartialUpdateOK with default headers values
func NewCircuitsCircuitTypesBulkPartialUpdateOK() *CircuitsCircuitTypesBulkPartialUpdateOK {
	return &CircuitsCircuitTypesBulkPartialUpdateOK{}
}

/* CircuitsCircuitTypesBulkPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTypesBulkPartialUpdateOK circuits circuit types bulk partial update o k
*/
type CircuitsCircuitTypesBulkPartialUpdateOK struct {
	Payload *models.CircuitType
}

func (o *CircuitsCircuitTypesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-types/][%d] circuitsCircuitTypesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTypesBulkPartialUpdateOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
