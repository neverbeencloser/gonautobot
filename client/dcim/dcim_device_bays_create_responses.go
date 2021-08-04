package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceBaysCreateReader is a Reader for the DcimDeviceBaysCreate structure.
type DcimDeviceBaysCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimDeviceBaysCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBaysCreateCreated creates a DcimDeviceBaysCreateCreated with default headers values
func NewDcimDeviceBaysCreateCreated() *DcimDeviceBaysCreateCreated {
	return &DcimDeviceBaysCreateCreated{}
}

/* DcimDeviceBaysCreateCreated describes a response with status code 201, with default header values.

DcimDeviceBaysCreateCreated dcim device bays create created
*/
type DcimDeviceBaysCreateCreated struct {
	Payload *models.DeviceBay
}

func (o *DcimDeviceBaysCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/device-bays/][%d] dcimDeviceBaysCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimDeviceBaysCreateCreated) GetPayload() *models.DeviceBay {
	return o.Payload
}

func (o *DcimDeviceBaysCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
