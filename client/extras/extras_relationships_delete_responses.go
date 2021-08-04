package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasRelationshipsDeleteReader is a Reader for the ExtrasRelationshipsDelete structure.
type ExtrasRelationshipsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasRelationshipsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasRelationshipsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasRelationshipsDeleteNoContent creates a ExtrasRelationshipsDeleteNoContent with default headers values
func NewExtrasRelationshipsDeleteNoContent() *ExtrasRelationshipsDeleteNoContent {
	return &ExtrasRelationshipsDeleteNoContent{}
}

/* ExtrasRelationshipsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasRelationshipsDeleteNoContent extras relationships delete no content
*/
type ExtrasRelationshipsDeleteNoContent struct {
}

func (o *ExtrasRelationshipsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/relationships/{id}/][%d] extrasRelationshipsDeleteNoContent ", 204)
}

func (o *ExtrasRelationshipsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
