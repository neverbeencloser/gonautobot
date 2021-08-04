package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceRolesCreateReader is a Reader for the DcimDeviceRolesCreate structure.
type DcimDeviceRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimDeviceRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceRolesCreateCreated creates a DcimDeviceRolesCreateCreated with default headers values
func NewDcimDeviceRolesCreateCreated() *DcimDeviceRolesCreateCreated {
	return &DcimDeviceRolesCreateCreated{}
}

/* DcimDeviceRolesCreateCreated describes a response with status code 201, with default header values.

DcimDeviceRolesCreateCreated dcim device roles create created
*/
type DcimDeviceRolesCreateCreated struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/device-roles/][%d] dcimDeviceRolesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimDeviceRolesCreateCreated) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
