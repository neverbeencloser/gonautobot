package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsProvidersPartialUpdateReader is a Reader for the CircuitsProvidersPartialUpdate structure.
type CircuitsProvidersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsProvidersPartialUpdateOK creates a CircuitsProvidersPartialUpdateOK with default headers values
func NewCircuitsProvidersPartialUpdateOK() *CircuitsProvidersPartialUpdateOK {
	return &CircuitsProvidersPartialUpdateOK{}
}

/* CircuitsProvidersPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsProvidersPartialUpdateOK circuits providers partial update o k
*/
type CircuitsProvidersPartialUpdateOK struct {
	Payload *models.Provider
}

func (o *CircuitsProvidersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/providers/{id}/][%d] circuitsProvidersPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsProvidersPartialUpdateOK) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
