package virtualization

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationVirtualMachinesDeleteReader is a Reader for the VirtualizationVirtualMachinesDelete structure.
type VirtualizationVirtualMachinesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationVirtualMachinesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesDeleteNoContent creates a VirtualizationVirtualMachinesDeleteNoContent with default headers values
func NewVirtualizationVirtualMachinesDeleteNoContent() *VirtualizationVirtualMachinesDeleteNoContent {
	return &VirtualizationVirtualMachinesDeleteNoContent{}
}

/* VirtualizationVirtualMachinesDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationVirtualMachinesDeleteNoContent virtualization virtual machines delete no content
*/
type VirtualizationVirtualMachinesDeleteNoContent struct {
}

func (o *VirtualizationVirtualMachinesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesDeleteNoContent ", 204)
}

func (o *VirtualizationVirtualMachinesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
