package circuits

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// CircuitsProvidersCreateReader is a Reader for the CircuitsProvidersCreate structure.
type CircuitsProvidersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCircuitsProvidersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsProvidersCreateCreated creates a CircuitsProvidersCreateCreated with default headers values
func NewCircuitsProvidersCreateCreated() *CircuitsProvidersCreateCreated {
	return &CircuitsProvidersCreateCreated{}
}

/* CircuitsProvidersCreateCreated describes a response with status code 201, with default header values.

CircuitsProvidersCreateCreated circuits providers create created
*/
type CircuitsProvidersCreateCreated struct {
	Payload *models.Provider
}

func (o *CircuitsProvidersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/providers/][%d] circuitsProvidersCreateCreated  %+v", 201, o.Payload)
}
func (o *CircuitsProvidersCreateCreated) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
