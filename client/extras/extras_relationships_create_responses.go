package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipsCreateReader is a Reader for the ExtrasRelationshipsCreate structure.
type ExtrasRelationshipsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasRelationshipsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipsCreateCreated creates a ExtrasRelationshipsCreateCreated with default headers values
func NewExtrasRelationshipsCreateCreated() *ExtrasRelationshipsCreateCreated {
	return &ExtrasRelationshipsCreateCreated{}
}

/* ExtrasRelationshipsCreateCreated describes a response with status code 201, with default header values.

ExtrasRelationshipsCreateCreated extras relationships create created
*/
type ExtrasRelationshipsCreateCreated struct {
	Payload *models.Relationship
}

func (o *ExtrasRelationshipsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/relationships/][%d] extrasRelationshipsCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasRelationshipsCreateCreated) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *ExtrasRelationshipsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
