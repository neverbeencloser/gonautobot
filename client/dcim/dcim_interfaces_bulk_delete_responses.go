package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimInterfacesBulkDeleteReader is a Reader for the DcimInterfacesBulkDelete structure.
type DcimInterfacesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInterfacesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfacesBulkDeleteNoContent creates a DcimInterfacesBulkDeleteNoContent with default headers values
func NewDcimInterfacesBulkDeleteNoContent() *DcimInterfacesBulkDeleteNoContent {
	return &DcimInterfacesBulkDeleteNoContent{}
}

/* DcimInterfacesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimInterfacesBulkDeleteNoContent dcim interfaces bulk delete no content
*/
type DcimInterfacesBulkDeleteNoContent struct {
}

func (o *DcimInterfacesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/interfaces/][%d] dcimInterfacesBulkDeleteNoContent ", 204)
}

func (o *DcimInterfacesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
