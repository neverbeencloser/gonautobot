package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasRelationshipAssociationsBulkDeleteReader is a Reader for the ExtrasRelationshipAssociationsBulkDelete structure.
type ExtrasRelationshipAssociationsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipAssociationsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasRelationshipAssociationsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipAssociationsBulkDeleteNoContent creates a ExtrasRelationshipAssociationsBulkDeleteNoContent with default headers values
func NewExtrasRelationshipAssociationsBulkDeleteNoContent() *ExtrasRelationshipAssociationsBulkDeleteNoContent {
	return &ExtrasRelationshipAssociationsBulkDeleteNoContent{}
}

/* ExtrasRelationshipAssociationsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasRelationshipAssociationsBulkDeleteNoContent extras relationship associations bulk delete no content
*/
type ExtrasRelationshipAssociationsBulkDeleteNoContent struct {
}

func (o *ExtrasRelationshipAssociationsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/relationship-associations/][%d] extrasRelationshipAssociationsBulkDeleteNoContent ", 204)
}

func (o *ExtrasRelationshipAssociationsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
