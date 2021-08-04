package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceRolesBulkPartialUpdateReader is a Reader for the DcimDeviceRolesBulkPartialUpdate structure.
type DcimDeviceRolesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceRolesBulkPartialUpdateOK creates a DcimDeviceRolesBulkPartialUpdateOK with default headers values
func NewDcimDeviceRolesBulkPartialUpdateOK() *DcimDeviceRolesBulkPartialUpdateOK {
	return &DcimDeviceRolesBulkPartialUpdateOK{}
}

/* DcimDeviceRolesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceRolesBulkPartialUpdateOK dcim device roles bulk partial update o k
*/
type DcimDeviceRolesBulkPartialUpdateOK struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-roles/][%d] dcimDeviceRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceRolesBulkPartialUpdateOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
