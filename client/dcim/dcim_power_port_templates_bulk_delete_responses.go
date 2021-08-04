package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerPortTemplatesBulkDeleteReader is a Reader for the DcimPowerPortTemplatesBulkDelete structure.
type DcimPowerPortTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPortTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesBulkDeleteNoContent creates a DcimPowerPortTemplatesBulkDeleteNoContent with default headers values
func NewDcimPowerPortTemplatesBulkDeleteNoContent() *DcimPowerPortTemplatesBulkDeleteNoContent {
	return &DcimPowerPortTemplatesBulkDeleteNoContent{}
}

/* DcimPowerPortTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPortTemplatesBulkDeleteNoContent dcim power port templates bulk delete no content
*/
type DcimPowerPortTemplatesBulkDeleteNoContent struct {
}

func (o *DcimPowerPortTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-port-templates/][%d] dcimPowerPortTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimPowerPortTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
