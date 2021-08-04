package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimConsoleServerPortTemplatesDeleteReader is a Reader for the DcimConsoleServerPortTemplatesDelete structure.
type DcimConsoleServerPortTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsoleServerPortTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesDeleteNoContent creates a DcimConsoleServerPortTemplatesDeleteNoContent with default headers values
func NewDcimConsoleServerPortTemplatesDeleteNoContent() *DcimConsoleServerPortTemplatesDeleteNoContent {
	return &DcimConsoleServerPortTemplatesDeleteNoContent{}
}

/* DcimConsoleServerPortTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimConsoleServerPortTemplatesDeleteNoContent dcim console server port templates delete no content
*/
type DcimConsoleServerPortTemplatesDeleteNoContent struct {
}

func (o *DcimConsoleServerPortTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimConsoleServerPortTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
