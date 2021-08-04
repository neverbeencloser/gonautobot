package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRearPortTemplatesBulkDeleteReader is a Reader for the DcimRearPortTemplatesBulkDelete structure.
type DcimRearPortTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRearPortTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortTemplatesBulkDeleteNoContent creates a DcimRearPortTemplatesBulkDeleteNoContent with default headers values
func NewDcimRearPortTemplatesBulkDeleteNoContent() *DcimRearPortTemplatesBulkDeleteNoContent {
	return &DcimRearPortTemplatesBulkDeleteNoContent{}
}

/* DcimRearPortTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRearPortTemplatesBulkDeleteNoContent dcim rear port templates bulk delete no content
*/
type DcimRearPortTemplatesBulkDeleteNoContent struct {
}

func (o *DcimRearPortTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rear-port-templates/][%d] dcimRearPortTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimRearPortTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
