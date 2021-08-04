package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipAssociationsBulkPartialUpdateReader is a Reader for the ExtrasRelationshipAssociationsBulkPartialUpdate structure.
type ExtrasRelationshipAssociationsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasRelationshipAssociationsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipAssociationsBulkPartialUpdateOK creates a ExtrasRelationshipAssociationsBulkPartialUpdateOK with default headers values
func NewExtrasRelationshipAssociationsBulkPartialUpdateOK() *ExtrasRelationshipAssociationsBulkPartialUpdateOK {
	return &ExtrasRelationshipAssociationsBulkPartialUpdateOK{}
}

/* ExtrasRelationshipAssociationsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasRelationshipAssociationsBulkPartialUpdateOK extras relationship associations bulk partial update o k
*/
type ExtrasRelationshipAssociationsBulkPartialUpdateOK struct {
	Payload *models.RelationshipAssociation
}

func (o *ExtrasRelationshipAssociationsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/relationship-associations/][%d] extrasRelationshipAssociationsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasRelationshipAssociationsBulkPartialUpdateOK) GetPayload() *models.RelationshipAssociation {
	return o.Payload
}

func (o *ExtrasRelationshipAssociationsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RelationshipAssociation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
