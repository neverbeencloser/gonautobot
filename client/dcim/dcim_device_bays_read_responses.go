package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceBaysReadReader is a Reader for the DcimDeviceBaysRead structure.
type DcimDeviceBaysReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBaysReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBaysReadOK creates a DcimDeviceBaysReadOK with default headers values
func NewDcimDeviceBaysReadOK() *DcimDeviceBaysReadOK {
	return &DcimDeviceBaysReadOK{}
}

/* DcimDeviceBaysReadOK describes a response with status code 200, with default header values.

DcimDeviceBaysReadOK dcim device bays read o k
*/
type DcimDeviceBaysReadOK struct {
	Payload *models.DeviceBay
}

func (o *DcimDeviceBaysReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/device-bays/{id}/][%d] dcimDeviceBaysReadOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceBaysReadOK) GetPayload() *models.DeviceBay {
	return o.Payload
}

func (o *DcimDeviceBaysReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
