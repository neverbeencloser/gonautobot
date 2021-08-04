package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationInterfacesBulkUpdateReader is a Reader for the VirtualizationInterfacesBulkUpdate structure.
type VirtualizationInterfacesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationInterfacesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationInterfacesBulkUpdateOK creates a VirtualizationInterfacesBulkUpdateOK with default headers values
func NewVirtualizationInterfacesBulkUpdateOK() *VirtualizationInterfacesBulkUpdateOK {
	return &VirtualizationInterfacesBulkUpdateOK{}
}

/* VirtualizationInterfacesBulkUpdateOK describes a response with status code 200, with default header values.

VirtualizationInterfacesBulkUpdateOK virtualization interfaces bulk update o k
*/
type VirtualizationInterfacesBulkUpdateOK struct {
	Payload *models.VMInterface
}

func (o *VirtualizationInterfacesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/interfaces/][%d] virtualizationInterfacesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationInterfacesBulkUpdateOK) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
