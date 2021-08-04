package virtualization

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// VirtualizationInterfacesPartialUpdateReader is a Reader for the VirtualizationInterfacesPartialUpdate structure.
type VirtualizationInterfacesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationInterfacesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationInterfacesPartialUpdateOK creates a VirtualizationInterfacesPartialUpdateOK with default headers values
func NewVirtualizationInterfacesPartialUpdateOK() *VirtualizationInterfacesPartialUpdateOK {
	return &VirtualizationInterfacesPartialUpdateOK{}
}

/* VirtualizationInterfacesPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationInterfacesPartialUpdateOK virtualization interfaces partial update o k
*/
type VirtualizationInterfacesPartialUpdateOK struct {
	Payload *models.VMInterface
}

func (o *VirtualizationInterfacesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/interfaces/{id}/][%d] virtualizationInterfacesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationInterfacesPartialUpdateOK) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
