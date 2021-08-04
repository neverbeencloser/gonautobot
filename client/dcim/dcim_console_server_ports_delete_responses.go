package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimConsoleServerPortsDeleteReader is a Reader for the DcimConsoleServerPortsDelete structure.
type DcimConsoleServerPortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsoleServerPortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortsDeleteNoContent creates a DcimConsoleServerPortsDeleteNoContent with default headers values
func NewDcimConsoleServerPortsDeleteNoContent() *DcimConsoleServerPortsDeleteNoContent {
	return &DcimConsoleServerPortsDeleteNoContent{}
}

/* DcimConsoleServerPortsDeleteNoContent describes a response with status code 204, with default header values.

DcimConsoleServerPortsDeleteNoContent dcim console server ports delete no content
*/
type DcimConsoleServerPortsDeleteNoContent struct {
}

func (o *DcimConsoleServerPortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsDeleteNoContent ", 204)
}

func (o *DcimConsoleServerPortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
