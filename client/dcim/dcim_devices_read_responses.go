package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDevicesReadReader is a Reader for the DcimDevicesRead structure.
type DcimDevicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDevicesReadOK creates a DcimDevicesReadOK with default headers values
func NewDcimDevicesReadOK() *DcimDevicesReadOK {
	return &DcimDevicesReadOK{}
}

/* DcimDevicesReadOK describes a response with status code 200, with default header values.

DcimDevicesReadOK dcim devices read o k
*/
type DcimDevicesReadOK struct {
	Payload *models.DeviceWithConfigContext
}

func (o *DcimDevicesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/devices/{id}/][%d] dcimDevicesReadOK  %+v", 200, o.Payload)
}
func (o *DcimDevicesReadOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
