package extras

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasCustomLinksDeleteReader is a Reader for the ExtrasCustomLinksDelete structure.
type ExtrasCustomLinksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomLinksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasCustomLinksDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomLinksDeleteNoContent creates a ExtrasCustomLinksDeleteNoContent with default headers values
func NewExtrasCustomLinksDeleteNoContent() *ExtrasCustomLinksDeleteNoContent {
	return &ExtrasCustomLinksDeleteNoContent{}
}

/* ExtrasCustomLinksDeleteNoContent describes a response with status code 204, with default header values.

ExtrasCustomLinksDeleteNoContent extras custom links delete no content
*/
type ExtrasCustomLinksDeleteNoContent struct {
}

func (o *ExtrasCustomLinksDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-links/{id}/][%d] extrasCustomLinksDeleteNoContent ", 204)
}

func (o *ExtrasCustomLinksDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
