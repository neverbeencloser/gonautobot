package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipAssociationsReadReader is a Reader for the ExtrasRelationshipAssociationsRead structure.
type ExtrasRelationshipAssociationsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipAssociationsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasRelationshipAssociationsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipAssociationsReadOK creates a ExtrasRelationshipAssociationsReadOK with default headers values
func NewExtrasRelationshipAssociationsReadOK() *ExtrasRelationshipAssociationsReadOK {
	return &ExtrasRelationshipAssociationsReadOK{}
}

/* ExtrasRelationshipAssociationsReadOK describes a response with status code 200, with default header values.

ExtrasRelationshipAssociationsReadOK extras relationship associations read o k
*/
type ExtrasRelationshipAssociationsReadOK struct {
	Payload *models.RelationshipAssociation
}

func (o *ExtrasRelationshipAssociationsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/relationship-associations/{id}/][%d] extrasRelationshipAssociationsReadOK  %+v", 200, o.Payload)
}
func (o *ExtrasRelationshipAssociationsReadOK) GetPayload() *models.RelationshipAssociation {
	return o.Payload
}

func (o *ExtrasRelationshipAssociationsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RelationshipAssociation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
