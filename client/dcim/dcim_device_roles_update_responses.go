package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceRolesUpdateReader is a Reader for the DcimDeviceRolesUpdate structure.
type DcimDeviceRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceRolesUpdateOK creates a DcimDeviceRolesUpdateOK with default headers values
func NewDcimDeviceRolesUpdateOK() *DcimDeviceRolesUpdateOK {
	return &DcimDeviceRolesUpdateOK{}
}

/* DcimDeviceRolesUpdateOK describes a response with status code 200, with default header values.

DcimDeviceRolesUpdateOK dcim device roles update o k
*/
type DcimDeviceRolesUpdateOK struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-roles/{id}/][%d] dcimDeviceRolesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceRolesUpdateOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
