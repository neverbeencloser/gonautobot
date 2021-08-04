package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipsReadReader is a Reader for the ExtrasRelationshipsRead structure.
type ExtrasRelationshipsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasRelationshipsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipsReadOK creates a ExtrasRelationshipsReadOK with default headers values
func NewExtrasRelationshipsReadOK() *ExtrasRelationshipsReadOK {
	return &ExtrasRelationshipsReadOK{}
}

/* ExtrasRelationshipsReadOK describes a response with status code 200, with default header values.

ExtrasRelationshipsReadOK extras relationships read o k
*/
type ExtrasRelationshipsReadOK struct {
	Payload *models.Relationship
}

func (o *ExtrasRelationshipsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/relationships/{id}/][%d] extrasRelationshipsReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasRelationshipsReadOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *ExtrasRelationshipsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
