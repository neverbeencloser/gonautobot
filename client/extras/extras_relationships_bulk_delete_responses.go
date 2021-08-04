package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasRelationshipsBulkDeleteReader is a Reader for the ExtrasRelationshipsBulkDelete structure.
type ExtrasRelationshipsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasRelationshipsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipsBulkDeleteNoContent creates a ExtrasRelationshipsBulkDeleteNoContent with default headers values
func NewExtrasRelationshipsBulkDeleteNoContent() *ExtrasRelationshipsBulkDeleteNoContent {
	return &ExtrasRelationshipsBulkDeleteNoContent{}
}

/* ExtrasRelationshipsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasRelationshipsBulkDeleteNoContent extras relationships bulk delete no content
*/
type ExtrasRelationshipsBulkDeleteNoContent struct {
}

func (o *ExtrasRelationshipsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/relationships/][%d] extrasRelationshipsBulkDeleteNoContent ", 204)
}

func (o *ExtrasRelationshipsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
