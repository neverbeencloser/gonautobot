package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationVirtualMachinesBulkDeleteReader is a Reader for the VirtualizationVirtualMachinesBulkDelete structure.
type VirtualizationVirtualMachinesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationVirtualMachinesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesBulkDeleteNoContent creates a VirtualizationVirtualMachinesBulkDeleteNoContent with default headers values
func NewVirtualizationVirtualMachinesBulkDeleteNoContent() *VirtualizationVirtualMachinesBulkDeleteNoContent {
	return &VirtualizationVirtualMachinesBulkDeleteNoContent{}
}

/* VirtualizationVirtualMachinesBulkDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationVirtualMachinesBulkDeleteNoContent virtualization virtual machines bulk delete no content
*/
type VirtualizationVirtualMachinesBulkDeleteNoContent struct {
}

func (o *VirtualizationVirtualMachinesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/virtual-machines/][%d] virtualizationVirtualMachinesBulkDeleteNoContent ", 204)
}

func (o *VirtualizationVirtualMachinesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
