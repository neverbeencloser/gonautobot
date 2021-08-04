package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceRolesBulkUpdateReader is a Reader for the DcimDeviceRolesBulkUpdate structure.
type DcimDeviceRolesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceRolesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceRolesBulkUpdateOK creates a DcimDeviceRolesBulkUpdateOK with default headers values
func NewDcimDeviceRolesBulkUpdateOK() *DcimDeviceRolesBulkUpdateOK {
	return &DcimDeviceRolesBulkUpdateOK{}
}

/* DcimDeviceRolesBulkUpdateOK describes a response with status code 200, with default header values.

DcimDeviceRolesBulkUpdateOK dcim device roles bulk update o k
*/
type DcimDeviceRolesBulkUpdateOK struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-roles/][%d] dcimDeviceRolesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceRolesBulkUpdateOK) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
