package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceTypesBulkPartialUpdateReader is a Reader for the DcimDeviceTypesBulkPartialUpdate structure.
type DcimDeviceTypesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceTypesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceTypesBulkPartialUpdateOK creates a DcimDeviceTypesBulkPartialUpdateOK with default headers values
func NewDcimDeviceTypesBulkPartialUpdateOK() *DcimDeviceTypesBulkPartialUpdateOK {
	return &DcimDeviceTypesBulkPartialUpdateOK{}
}

/* DcimDeviceTypesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceTypesBulkPartialUpdateOK dcim device types bulk partial update o k
*/
type DcimDeviceTypesBulkPartialUpdateOK struct {
	Payload *models.DeviceType
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcimDeviceTypesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceTypesBulkPartialUpdateOK) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
