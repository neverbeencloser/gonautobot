package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceRolesBulkDeleteReader is a Reader for the DcimDeviceRolesBulkDelete structure.
type DcimDeviceRolesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceRolesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceRolesBulkDeleteNoContent creates a DcimDeviceRolesBulkDeleteNoContent with default headers values
func NewDcimDeviceRolesBulkDeleteNoContent() *DcimDeviceRolesBulkDeleteNoContent {
	return &DcimDeviceRolesBulkDeleteNoContent{}
}

/* DcimDeviceRolesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceRolesBulkDeleteNoContent dcim device roles bulk delete no content
*/
type DcimDeviceRolesBulkDeleteNoContent struct {
}

func (o *DcimDeviceRolesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/][%d] dcimDeviceRolesBulkDeleteNoContent ", 204)
}

func (o *DcimDeviceRolesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
