package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipsBulkPartialUpdateReader is a Reader for the ExtrasRelationshipsBulkPartialUpdate structure.
type ExtrasRelationshipsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasRelationshipsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipsBulkPartialUpdateOK creates a ExtrasRelationshipsBulkPartialUpdateOK with default headers values
func NewExtrasRelationshipsBulkPartialUpdateOK() *ExtrasRelationshipsBulkPartialUpdateOK {
	return &ExtrasRelationshipsBulkPartialUpdateOK{}
}

/* ExtrasRelationshipsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasRelationshipsBulkPartialUpdateOK extras relationships bulk partial update o k
*/
type ExtrasRelationshipsBulkPartialUpdateOK struct {
	Payload *models.Relationship
}

func (o *ExtrasRelationshipsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/relationships/][%d] extrasRelationshipsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasRelationshipsBulkPartialUpdateOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *ExtrasRelationshipsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
