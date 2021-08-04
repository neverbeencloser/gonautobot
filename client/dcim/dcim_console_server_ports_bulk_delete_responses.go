package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimConsoleServerPortsBulkDeleteReader is a Reader for the DcimConsoleServerPortsBulkDelete structure.
type DcimConsoleServerPortsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsoleServerPortsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortsBulkDeleteNoContent creates a DcimConsoleServerPortsBulkDeleteNoContent with default headers values
func NewDcimConsoleServerPortsBulkDeleteNoContent() *DcimConsoleServerPortsBulkDeleteNoContent {
	return &DcimConsoleServerPortsBulkDeleteNoContent{}
}

/* DcimConsoleServerPortsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimConsoleServerPortsBulkDeleteNoContent dcim console server ports bulk delete no content
*/
type DcimConsoleServerPortsBulkDeleteNoContent struct {
}

func (o *DcimConsoleServerPortsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-ports/][%d] dcimConsoleServerPortsBulkDeleteNoContent ", 204)
}

func (o *DcimConsoleServerPortsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
