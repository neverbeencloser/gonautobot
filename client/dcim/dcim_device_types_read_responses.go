package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceTypesReadReader is a Reader for the DcimDeviceTypesRead structure.
type DcimDeviceTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceTypesReadOK creates a DcimDeviceTypesReadOK with default headers values
func NewDcimDeviceTypesReadOK() *DcimDeviceTypesReadOK {
	return &DcimDeviceTypesReadOK{}
}

/* DcimDeviceTypesReadOK describes a response with status code 200, with default header values.

DcimDeviceTypesReadOK dcim device types read o k
*/
type DcimDeviceTypesReadOK struct {
	Payload *models.DeviceType
}

func (o *DcimDeviceTypesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/device-types/{id}/][%d] dcimDeviceTypesReadOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceTypesReadOK) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
