package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimFrontPortTemplatesBulkDeleteReader is a Reader for the DcimFrontPortTemplatesBulkDelete structure.
type DcimFrontPortTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimFrontPortTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortTemplatesBulkDeleteNoContent creates a DcimFrontPortTemplatesBulkDeleteNoContent with default headers values
func NewDcimFrontPortTemplatesBulkDeleteNoContent() *DcimFrontPortTemplatesBulkDeleteNoContent {
	return &DcimFrontPortTemplatesBulkDeleteNoContent{}
}

/* DcimFrontPortTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimFrontPortTemplatesBulkDeleteNoContent dcim front port templates bulk delete no content
*/
type DcimFrontPortTemplatesBulkDeleteNoContent struct {
}

func (o *DcimFrontPortTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/front-port-templates/][%d] dcimFrontPortTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimFrontPortTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
