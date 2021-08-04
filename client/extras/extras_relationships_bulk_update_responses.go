package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipsBulkUpdateReader is a Reader for the ExtrasRelationshipsBulkUpdate structure.
type ExtrasRelationshipsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasRelationshipsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipsBulkUpdateOK creates a ExtrasRelationshipsBulkUpdateOK with default headers values
func NewExtrasRelationshipsBulkUpdateOK() *ExtrasRelationshipsBulkUpdateOK {
	return &ExtrasRelationshipsBulkUpdateOK{}
}

/* ExtrasRelationshipsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasRelationshipsBulkUpdateOK extras relationships bulk update o k
*/
type ExtrasRelationshipsBulkUpdateOK struct {
	Payload *models.Relationship
}

func (o *ExtrasRelationshipsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/relationships/][%d] extrasRelationshipsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasRelationshipsBulkUpdateOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *ExtrasRelationshipsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
