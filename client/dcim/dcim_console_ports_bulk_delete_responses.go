package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimConsolePortsBulkDeleteReader is a Reader for the DcimConsolePortsBulkDelete structure.
type DcimConsolePortsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsolePortsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortsBulkDeleteNoContent creates a DcimConsolePortsBulkDeleteNoContent with default headers values
func NewDcimConsolePortsBulkDeleteNoContent() *DcimConsolePortsBulkDeleteNoContent {
	return &DcimConsolePortsBulkDeleteNoContent{}
}

/* DcimConsolePortsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimConsolePortsBulkDeleteNoContent dcim console ports bulk delete no content
*/
type DcimConsolePortsBulkDeleteNoContent struct {
}

func (o *DcimConsolePortsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-ports/][%d] dcimConsolePortsBulkDeleteNoContent ", 204)
}

func (o *DcimConsolePortsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
