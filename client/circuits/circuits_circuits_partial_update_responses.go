package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsCircuitsPartialUpdateReader is a Reader for the CircuitsCircuitsPartialUpdate structure.
type CircuitsCircuitsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitsPartialUpdateOK creates a CircuitsCircuitsPartialUpdateOK with default headers values
func NewCircuitsCircuitsPartialUpdateOK() *CircuitsCircuitsPartialUpdateOK {
	return &CircuitsCircuitsPartialUpdateOK{}
}

/* CircuitsCircuitsPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitsPartialUpdateOK circuits circuits partial update o k
*/
type CircuitsCircuitsPartialUpdateOK struct {
	Payload *models.Circuit
}

func (o *CircuitsCircuitsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuits/{id}/][%d] circuitsCircuitsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitsPartialUpdateOK) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
