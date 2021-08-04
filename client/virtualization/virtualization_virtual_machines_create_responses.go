package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationVirtualMachinesCreateReader is a Reader for the VirtualizationVirtualMachinesCreate structure.
type VirtualizationVirtualMachinesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVirtualizationVirtualMachinesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesCreateCreated creates a VirtualizationVirtualMachinesCreateCreated with default headers values
func NewVirtualizationVirtualMachinesCreateCreated() *VirtualizationVirtualMachinesCreateCreated {
	return &VirtualizationVirtualMachinesCreateCreated{}
}

/* VirtualizationVirtualMachinesCreateCreated describes a response with status code 201, with default header values.

VirtualizationVirtualMachinesCreateCreated virtualization virtual machines create created
*/
type VirtualizationVirtualMachinesCreateCreated struct {
	Payload *models.VirtualMachineWithConfigContext
}

func (o *VirtualizationVirtualMachinesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /virtualization/virtual-machines/][%d] virtualizationVirtualMachinesCreateCreated  %+v", 201, o.Payload)
}
func (o *VirtualizationVirtualMachinesCreateCreated) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
