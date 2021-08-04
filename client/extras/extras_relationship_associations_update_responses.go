package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipAssociationsUpdateReader is a Reader for the ExtrasRelationshipAssociationsUpdate structure.
type ExtrasRelationshipAssociationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipAssociationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasRelationshipAssociationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipAssociationsUpdateOK creates a ExtrasRelationshipAssociationsUpdateOK with default headers values
func NewExtrasRelationshipAssociationsUpdateOK() *ExtrasRelationshipAssociationsUpdateOK {
	return &ExtrasRelationshipAssociationsUpdateOK{}
}

/* ExtrasRelationshipAssociationsUpdateOK describes a response with status code 200, with default header values.

ExtrasRelationshipAssociationsUpdateOK extras relationship associations update o k
*/
type ExtrasRelationshipAssociationsUpdateOK struct {
	Payload *models.RelationshipAssociation
}

func (o *ExtrasRelationshipAssociationsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/relationship-associations/{id}/][%d] extrasRelationshipAssociationsUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasRelationshipAssociationsUpdateOK) GetPayload() *models.RelationshipAssociation {
	return o.Payload
}

func (o *ExtrasRelationshipAssociationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RelationshipAssociation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
