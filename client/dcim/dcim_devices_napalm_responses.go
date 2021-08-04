package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDevicesNapalmReader is a Reader for the DcimDevicesNapalm structure.
type DcimDevicesNapalmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesNapalmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesNapalmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDevicesNapalmOK creates a DcimDevicesNapalmOK with default headers values
func NewDcimDevicesNapalmOK() *DcimDevicesNapalmOK {
	return &DcimDevicesNapalmOK{}
}

/* DcimDevicesNapalmOK describes a response with status code 200, with default header values.

DcimDevicesNapalmOK dcim devices napalm o k
*/
type DcimDevicesNapalmOK struct {
	Payload *models.DeviceNAPALM
}

func (o *DcimDevicesNapalmOK) Error() string {
	return fmt.Sprintf("[GET /dcim/devices/{id}/napalm/][%d] dcimDevicesNapalmOK  %+v", 200, o.Payload)
}
func (o *DcimDevicesNapalmOK) GetPayload() *models.DeviceNAPALM {
	return o.Payload
}

func (o *DcimDevicesNapalmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceNAPALM)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
