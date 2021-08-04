package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationVirtualMachinesBulkPartialUpdateReader is a Reader for the VirtualizationVirtualMachinesBulkPartialUpdate structure.
type VirtualizationVirtualMachinesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationVirtualMachinesBulkPartialUpdateOK creates a VirtualizationVirtualMachinesBulkPartialUpdateOK with default headers values
func NewVirtualizationVirtualMachinesBulkPartialUpdateOK() *VirtualizationVirtualMachinesBulkPartialUpdateOK {
	return &VirtualizationVirtualMachinesBulkPartialUpdateOK{}
}

/* VirtualizationVirtualMachinesBulkPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesBulkPartialUpdateOK virtualization virtual machines bulk partial update o k
*/
type VirtualizationVirtualMachinesBulkPartialUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

func (o *VirtualizationVirtualMachinesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/virtual-machines/][%d] virtualizationVirtualMachinesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationVirtualMachinesBulkPartialUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
