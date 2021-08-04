package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceTypesBulkDeleteReader is a Reader for the DcimDeviceTypesBulkDelete structure.
type DcimDeviceTypesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceTypesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceTypesBulkDeleteNoContent creates a DcimDeviceTypesBulkDeleteNoContent with default headers values
func NewDcimDeviceTypesBulkDeleteNoContent() *DcimDeviceTypesBulkDeleteNoContent {
	return &DcimDeviceTypesBulkDeleteNoContent{}
}

/* DcimDeviceTypesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceTypesBulkDeleteNoContent dcim device types bulk delete no content
*/
type DcimDeviceTypesBulkDeleteNoContent struct {
}

func (o *DcimDeviceTypesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/][%d] dcimDeviceTypesBulkDeleteNoContent ", 204)
}

func (o *DcimDeviceTypesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
