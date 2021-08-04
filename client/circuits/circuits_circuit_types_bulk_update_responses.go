package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitTypesBulkUpdateReader is a Reader for the CircuitsCircuitTypesBulkUpdate structure.
type CircuitsCircuitTypesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTypesBulkUpdateOK creates a CircuitsCircuitTypesBulkUpdateOK with default headers values
func NewCircuitsCircuitTypesBulkUpdateOK() *CircuitsCircuitTypesBulkUpdateOK {
	return &CircuitsCircuitTypesBulkUpdateOK{}
}

/* CircuitsCircuitTypesBulkUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTypesBulkUpdateOK circuits circuit types bulk update o k
*/
type CircuitsCircuitTypesBulkUpdateOK struct {
	Payload *models.CircuitType
}

func (o *CircuitsCircuitTypesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/][%d] circuitsCircuitTypesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTypesBulkUpdateOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
