package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDevicesBulkPartialUpdateReader is a Reader for the DcimDevicesBulkPartialUpdate structure.
type DcimDevicesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDevicesBulkPartialUpdateOK creates a DcimDevicesBulkPartialUpdateOK with default headers values
func NewDcimDevicesBulkPartialUpdateOK() *DcimDevicesBulkPartialUpdateOK {
	return &DcimDevicesBulkPartialUpdateOK{}
}

/* DcimDevicesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDevicesBulkPartialUpdateOK dcim devices bulk partial update o k
*/
type DcimDevicesBulkPartialUpdateOK struct {
	Payload *models.DeviceWithConfigContext
}

func (o *DcimDevicesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/devices/][%d] dcimDevicesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDevicesBulkPartialUpdateOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
