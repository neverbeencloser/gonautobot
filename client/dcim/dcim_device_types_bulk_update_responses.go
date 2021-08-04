package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceTypesBulkUpdateReader is a Reader for the DcimDeviceTypesBulkUpdate structure.
type DcimDeviceTypesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceTypesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceTypesBulkUpdateOK creates a DcimDeviceTypesBulkUpdateOK with default headers values
func NewDcimDeviceTypesBulkUpdateOK() *DcimDeviceTypesBulkUpdateOK {
	return &DcimDeviceTypesBulkUpdateOK{}
}

/* DcimDeviceTypesBulkUpdateOK describes a response with status code 200, with default header values.

DcimDeviceTypesBulkUpdateOK dcim device types bulk update o k
*/
type DcimDeviceTypesBulkUpdateOK struct {
	Payload *models.DeviceType
}

func (o *DcimDeviceTypesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-types/][%d] dcimDeviceTypesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceTypesBulkUpdateOK) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
