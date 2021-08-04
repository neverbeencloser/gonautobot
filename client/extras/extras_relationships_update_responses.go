package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipsUpdateReader is a Reader for the ExtrasRelationshipsUpdate structure.
type ExtrasRelationshipsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasRelationshipsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipsUpdateOK creates a ExtrasRelationshipsUpdateOK with default headers values
func NewExtrasRelationshipsUpdateOK() *ExtrasRelationshipsUpdateOK {
	return &ExtrasRelationshipsUpdateOK{}
}

/* ExtrasRelationshipsUpdateOK describes a response with status code 200, with default header values.

ExtrasRelationshipsUpdateOK extras relationships update o k
*/
type ExtrasRelationshipsUpdateOK struct {
	Payload *models.Relationship
}

func (o *ExtrasRelationshipsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/relationships/{id}/][%d] extrasRelationshipsUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasRelationshipsUpdateOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *ExtrasRelationshipsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
