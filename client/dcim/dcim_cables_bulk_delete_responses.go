package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimCablesBulkDeleteReader is a Reader for the DcimCablesBulkDelete structure.
type DcimCablesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimCablesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimCablesBulkDeleteNoContent creates a DcimCablesBulkDeleteNoContent with default headers values
func NewDcimCablesBulkDeleteNoContent() *DcimCablesBulkDeleteNoContent {
	return &DcimCablesBulkDeleteNoContent{}
}

/* DcimCablesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimCablesBulkDeleteNoContent dcim cables bulk delete no content
*/
type DcimCablesBulkDeleteNoContent struct {
}

func (o *DcimCablesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cables/][%d] dcimCablesBulkDeleteNoContent ", 204)
}

func (o *DcimCablesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
