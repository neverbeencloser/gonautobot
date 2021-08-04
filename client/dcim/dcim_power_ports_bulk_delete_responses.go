package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerPortsBulkDeleteReader is a Reader for the DcimPowerPortsBulkDelete structure.
type DcimPowerPortsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPortsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortsBulkDeleteNoContent creates a DcimPowerPortsBulkDeleteNoContent with default headers values
func NewDcimPowerPortsBulkDeleteNoContent() *DcimPowerPortsBulkDeleteNoContent {
	return &DcimPowerPortsBulkDeleteNoContent{}
}

/* DcimPowerPortsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPortsBulkDeleteNoContent dcim power ports bulk delete no content
*/
type DcimPowerPortsBulkDeleteNoContent struct {
}

func (o *DcimPowerPortsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-ports/][%d] dcimPowerPortsBulkDeleteNoContent ", 204)
}

func (o *DcimPowerPortsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
