package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimConsoleServerPortTemplatesBulkDeleteReader is a Reader for the DcimConsoleServerPortTemplatesBulkDelete structure.
type DcimConsoleServerPortTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsoleServerPortTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesBulkDeleteNoContent creates a DcimConsoleServerPortTemplatesBulkDeleteNoContent with default headers values
func NewDcimConsoleServerPortTemplatesBulkDeleteNoContent() *DcimConsoleServerPortTemplatesBulkDeleteNoContent {
	return &DcimConsoleServerPortTemplatesBulkDeleteNoContent{}
}

/* DcimConsoleServerPortTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimConsoleServerPortTemplatesBulkDeleteNoContent dcim console server port templates bulk delete no content
*/
type DcimConsoleServerPortTemplatesBulkDeleteNoContent struct {
}

func (o *DcimConsoleServerPortTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-port-templates/][%d] dcimConsoleServerPortTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimConsoleServerPortTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
