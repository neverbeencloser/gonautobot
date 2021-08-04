package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipAssociationsPartialUpdateReader is a Reader for the ExtrasRelationshipAssociationsPartialUpdate structure.
type ExtrasRelationshipAssociationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipAssociationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasRelationshipAssociationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipAssociationsPartialUpdateOK creates a ExtrasRelationshipAssociationsPartialUpdateOK with default headers values
func NewExtrasRelationshipAssociationsPartialUpdateOK() *ExtrasRelationshipAssociationsPartialUpdateOK {
	return &ExtrasRelationshipAssociationsPartialUpdateOK{}
}

/* ExtrasRelationshipAssociationsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasRelationshipAssociationsPartialUpdateOK extras relationship associations partial update o k
*/
type ExtrasRelationshipAssociationsPartialUpdateOK struct {
	Payload *models.RelationshipAssociation
}

func (o *ExtrasRelationshipAssociationsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/relationship-associations/{id}/][%d] extrasRelationshipAssociationsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasRelationshipAssociationsPartialUpdateOK) GetPayload() *models.RelationshipAssociation {
	return o.Payload
}

func (o *ExtrasRelationshipAssociationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RelationshipAssociation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
