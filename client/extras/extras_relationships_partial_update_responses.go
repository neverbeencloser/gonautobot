package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipsPartialUpdateReader is a Reader for the ExtrasRelationshipsPartialUpdate structure.
type ExtrasRelationshipsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasRelationshipsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipsPartialUpdateOK creates a ExtrasRelationshipsPartialUpdateOK with default headers values
func NewExtrasRelationshipsPartialUpdateOK() *ExtrasRelationshipsPartialUpdateOK {
	return &ExtrasRelationshipsPartialUpdateOK{}
}

/* ExtrasRelationshipsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasRelationshipsPartialUpdateOK extras relationships partial update o k
*/
type ExtrasRelationshipsPartialUpdateOK struct {
	Payload *models.Relationship
}

func (o *ExtrasRelationshipsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/relationships/{id}/][%d] extrasRelationshipsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasRelationshipsPartialUpdateOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *ExtrasRelationshipsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
