package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipAssociationsBulkUpdateReader is a Reader for the ExtrasRelationshipAssociationsBulkUpdate structure.
type ExtrasRelationshipAssociationsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipAssociationsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasRelationshipAssociationsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipAssociationsBulkUpdateOK creates a ExtrasRelationshipAssociationsBulkUpdateOK with default headers values
func NewExtrasRelationshipAssociationsBulkUpdateOK() *ExtrasRelationshipAssociationsBulkUpdateOK {
	return &ExtrasRelationshipAssociationsBulkUpdateOK{}
}

/* ExtrasRelationshipAssociationsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasRelationshipAssociationsBulkUpdateOK extras relationship associations bulk update o k
*/
type ExtrasRelationshipAssociationsBulkUpdateOK struct {
	Payload *models.RelationshipAssociation
}

func (o *ExtrasRelationshipAssociationsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/relationship-associations/][%d] extrasRelationshipAssociationsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *ExtrasRelationshipAssociationsBulkUpdateOK) GetPayload() *models.RelationshipAssociation {
	return o.Payload
}

func (o *ExtrasRelationshipAssociationsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RelationshipAssociation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
