package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasStatusesDeleteReader is a Reader for the ExtrasStatusesDelete structure.
type ExtrasStatusesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasStatusesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasStatusesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasStatusesDeleteNoContent creates a ExtrasStatusesDeleteNoContent with default headers values
func NewExtrasStatusesDeleteNoContent() *ExtrasStatusesDeleteNoContent {
	return &ExtrasStatusesDeleteNoContent{}
}

/* ExtrasStatusesDeleteNoContent describes a response with status code 204, with default header values.

ExtrasStatusesDeleteNoContent extras statuses delete no content
*/
type ExtrasStatusesDeleteNoContent struct {
}

func (o *ExtrasStatusesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/statuses/{id}/][%d] extrasStatusesDeleteNoContent ", 204)
}

func (o *ExtrasStatusesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
