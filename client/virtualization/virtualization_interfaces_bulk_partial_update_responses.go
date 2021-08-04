package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationInterfacesBulkPartialUpdateReader is a Reader for the VirtualizationInterfacesBulkPartialUpdate structure.
type VirtualizationInterfacesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationInterfacesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationInterfacesBulkPartialUpdateOK creates a VirtualizationInterfacesBulkPartialUpdateOK with default headers values
func NewVirtualizationInterfacesBulkPartialUpdateOK() *VirtualizationInterfacesBulkPartialUpdateOK {
	return &VirtualizationInterfacesBulkPartialUpdateOK{}
}

/* VirtualizationInterfacesBulkPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationInterfacesBulkPartialUpdateOK virtualization interfaces bulk partial update o k
*/
type VirtualizationInterfacesBulkPartialUpdateOK struct {
	Payload *models.VMInterface
}

func (o *VirtualizationInterfacesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/interfaces/][%d] virtualizationInterfacesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationInterfacesBulkPartialUpdateOK) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
