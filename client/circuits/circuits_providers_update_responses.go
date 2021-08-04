package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsProvidersUpdateReader is a Reader for the CircuitsProvidersUpdate structure.
type CircuitsProvidersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsProvidersUpdateOK creates a CircuitsProvidersUpdateOK with default headers values
func NewCircuitsProvidersUpdateOK() *CircuitsProvidersUpdateOK {
	return &CircuitsProvidersUpdateOK{}
}

/* CircuitsProvidersUpdateOK describes a response with status code 200, with default header values.

CircuitsProvidersUpdateOK circuits providers update o k
*/
type CircuitsProvidersUpdateOK struct {
	Payload *models.Provider
}

func (o *CircuitsProvidersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/providers/{id}/][%d] circuitsProvidersUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsProvidersUpdateOK) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
