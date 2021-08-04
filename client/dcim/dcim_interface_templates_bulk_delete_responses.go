package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimInterfaceTemplatesBulkDeleteReader is a Reader for the DcimInterfaceTemplatesBulkDelete structure.
type DcimInterfaceTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInterfaceTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfaceTemplatesBulkDeleteNoContent creates a DcimInterfaceTemplatesBulkDeleteNoContent with default headers values
func NewDcimInterfaceTemplatesBulkDeleteNoContent() *DcimInterfaceTemplatesBulkDeleteNoContent {
	return &DcimInterfaceTemplatesBulkDeleteNoContent{}
}

/* DcimInterfaceTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimInterfaceTemplatesBulkDeleteNoContent dcim interface templates bulk delete no content
*/
type DcimInterfaceTemplatesBulkDeleteNoContent struct {
}

func (o *DcimInterfaceTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/interface-templates/][%d] dcimInterfaceTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimInterfaceTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
