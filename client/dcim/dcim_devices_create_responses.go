package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDevicesCreateReader is a Reader for the DcimDevicesCreate structure.
type DcimDevicesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimDevicesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDevicesCreateCreated creates a DcimDevicesCreateCreated with default headers values
func NewDcimDevicesCreateCreated() *DcimDevicesCreateCreated {
	return &DcimDevicesCreateCreated{}
}

/* DcimDevicesCreateCreated describes a response with status code 201, with default header values.

DcimDevicesCreateCreated dcim devices create created
*/
type DcimDevicesCreateCreated struct {
	Payload *models.DeviceWithConfigContext
}

func (o *DcimDevicesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/devices/][%d] dcimDevicesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimDevicesCreateCreated) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
