package extras

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// ExtrasRelationshipAssociationsCreateReader is a Reader for the ExtrasRelationshipAssociationsCreate structure.
type ExtrasRelationshipAssociationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipAssociationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasRelationshipAssociationsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipAssociationsCreateCreated creates a ExtrasRelationshipAssociationsCreateCreated with default headers values
func NewExtrasRelationshipAssociationsCreateCreated() *ExtrasRelationshipAssociationsCreateCreated {
	return &ExtrasRelationshipAssociationsCreateCreated{}
}

/* ExtrasRelationshipAssociationsCreateCreated describes a response with status code 201, with default header values.

ExtrasRelationshipAssociationsCreateCreated extras relationship associations create created
*/
type ExtrasRelationshipAssociationsCreateCreated struct {
	Payload *models.RelationshipAssociation
}

func (o *ExtrasRelationshipAssociationsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/relationship-associations/][%d] extrasRelationshipAssociationsCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasRelationshipAssociationsCreateCreated) GetPayload() *models.RelationshipAssociation {
	return o.Payload
}

func (o *ExtrasRelationshipAssociationsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RelationshipAssociation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
