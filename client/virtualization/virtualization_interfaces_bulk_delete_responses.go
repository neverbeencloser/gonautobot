package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationInterfacesBulkDeleteReader is a Reader for the VirtualizationInterfacesBulkDelete structure.
type VirtualizationInterfacesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationInterfacesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationInterfacesBulkDeleteNoContent creates a VirtualizationInterfacesBulkDeleteNoContent with default headers values
func NewVirtualizationInterfacesBulkDeleteNoContent() *VirtualizationInterfacesBulkDeleteNoContent {
	return &VirtualizationInterfacesBulkDeleteNoContent{}
}

/* VirtualizationInterfacesBulkDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationInterfacesBulkDeleteNoContent virtualization interfaces bulk delete no content
*/
type VirtualizationInterfacesBulkDeleteNoContent struct {
}

func (o *VirtualizationInterfacesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/interfaces/][%d] virtualizationInterfacesBulkDeleteNoContent ", 204)
}

func (o *VirtualizationInterfacesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
