package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceBaysPartialUpdateReader is a Reader for the DcimDeviceBaysPartialUpdate structure.
type DcimDeviceBaysPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBaysPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBaysPartialUpdateOK creates a DcimDeviceBaysPartialUpdateOK with default headers values
func NewDcimDeviceBaysPartialUpdateOK() *DcimDeviceBaysPartialUpdateOK {
	return &DcimDeviceBaysPartialUpdateOK{}
}

/* DcimDeviceBaysPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBaysPartialUpdateOK dcim device bays partial update o k
*/
type DcimDeviceBaysPartialUpdateOK struct {
	Payload *models.DeviceBay
}

func (o *DcimDeviceBaysPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bays/{id}/][%d] dcimDeviceBaysPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceBaysPartialUpdateOK) GetPayload() *models.DeviceBay {
	return o.Payload
}

func (o *DcimDeviceBaysPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
