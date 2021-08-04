package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationInterfacesCreateReader is a Reader for the VirtualizationInterfacesCreate structure.
type VirtualizationInterfacesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVirtualizationInterfacesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationInterfacesCreateCreated creates a VirtualizationInterfacesCreateCreated with default headers values
func NewVirtualizationInterfacesCreateCreated() *VirtualizationInterfacesCreateCreated {
	return &VirtualizationInterfacesCreateCreated{}
}

/* VirtualizationInterfacesCreateCreated describes a response with status code 201, with default header values.

VirtualizationInterfacesCreateCreated virtualization interfaces create created
*/
type VirtualizationInterfacesCreateCreated struct {
	Payload *models.VMInterface
}

func (o *VirtualizationInterfacesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /virtualization/interfaces/][%d] virtualizationInterfacesCreateCreated  %+v", 201, o.Payload)
}
func (o *VirtualizationInterfacesCreateCreated) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
