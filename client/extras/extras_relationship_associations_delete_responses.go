package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasRelationshipAssociationsDeleteReader is a Reader for the ExtrasRelationshipAssociationsDelete structure.
type ExtrasRelationshipAssociationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipAssociationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasRelationshipAssociationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipAssociationsDeleteNoContent creates a ExtrasRelationshipAssociationsDeleteNoContent with default headers values
func NewExtrasRelationshipAssociationsDeleteNoContent() *ExtrasRelationshipAssociationsDeleteNoContent {
	return &ExtrasRelationshipAssociationsDeleteNoContent{}
}

/* ExtrasRelationshipAssociationsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasRelationshipAssociationsDeleteNoContent extras relationship associations delete no content
*/
type ExtrasRelationshipAssociationsDeleteNoContent struct {
}

func (o *ExtrasRelationshipAssociationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/relationship-associations/{id}/][%d] extrasRelationshipAssociationsDeleteNoContent ", 204)
}

func (o *ExtrasRelationshipAssociationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
