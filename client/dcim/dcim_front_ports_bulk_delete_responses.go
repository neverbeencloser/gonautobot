package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimFrontPortsBulkDeleteReader is a Reader for the DcimFrontPortsBulkDelete structure.
type DcimFrontPortsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimFrontPortsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortsBulkDeleteNoContent creates a DcimFrontPortsBulkDeleteNoContent with default headers values
func NewDcimFrontPortsBulkDeleteNoContent() *DcimFrontPortsBulkDeleteNoContent {
	return &DcimFrontPortsBulkDeleteNoContent{}
}

/* DcimFrontPortsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimFrontPortsBulkDeleteNoContent dcim front ports bulk delete no content
*/
type DcimFrontPortsBulkDeleteNoContent struct {
}

func (o *DcimFrontPortsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/front-ports/][%d] dcimFrontPortsBulkDeleteNoContent ", 204)
}

func (o *DcimFrontPortsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
