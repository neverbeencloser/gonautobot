package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasTagsDeleteReader is a Reader for the ExtrasTagsDelete structure.
type ExtrasTagsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasTagsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasTagsDeleteNoContent creates a ExtrasTagsDeleteNoContent with default headers values
func NewExtrasTagsDeleteNoContent() *ExtrasTagsDeleteNoContent {
	return &ExtrasTagsDeleteNoContent{}
}

/* ExtrasTagsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasTagsDeleteNoContent extras tags delete no content
*/
type ExtrasTagsDeleteNoContent struct {
}

func (o *ExtrasTagsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/tags/{id}/][%d] extrasTagsDeleteNoContent ", 204)
}

func (o *ExtrasTagsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
