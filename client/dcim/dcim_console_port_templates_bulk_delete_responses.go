package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimConsolePortTemplatesBulkDeleteReader is a Reader for the DcimConsolePortTemplatesBulkDelete structure.
type DcimConsolePortTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsolePortTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimConsolePortTemplatesBulkDeleteNoContent creates a DcimConsolePortTemplatesBulkDeleteNoContent with default headers values
func NewDcimConsolePortTemplatesBulkDeleteNoContent() *DcimConsolePortTemplatesBulkDeleteNoContent {
	return &DcimConsolePortTemplatesBulkDeleteNoContent{}
}

/* DcimConsolePortTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimConsolePortTemplatesBulkDeleteNoContent dcim console port templates bulk delete no content
*/
type DcimConsolePortTemplatesBulkDeleteNoContent struct {
}

func (o *DcimConsolePortTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-port-templates/][%d] dcimConsolePortTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimConsolePortTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
