package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerPanelsBulkDeleteReader is a Reader for the DcimPowerPanelsBulkDelete structure.
type DcimPowerPanelsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPanelsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPanelsBulkDeleteNoContent creates a DcimPowerPanelsBulkDeleteNoContent with default headers values
func NewDcimPowerPanelsBulkDeleteNoContent() *DcimPowerPanelsBulkDeleteNoContent {
	return &DcimPowerPanelsBulkDeleteNoContent{}
}

/* DcimPowerPanelsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPanelsBulkDeleteNoContent dcim power panels bulk delete no content
*/
type DcimPowerPanelsBulkDeleteNoContent struct {
}

func (o *DcimPowerPanelsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-panels/][%d] dcimPowerPanelsBulkDeleteNoContent ", 204)
}

func (o *DcimPowerPanelsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
