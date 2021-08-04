package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceBaysBulkPartialUpdateReader is a Reader for the DcimDeviceBaysBulkPartialUpdate structure.
type DcimDeviceBaysBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBaysBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBaysBulkPartialUpdateOK creates a DcimDeviceBaysBulkPartialUpdateOK with default headers values
func NewDcimDeviceBaysBulkPartialUpdateOK() *DcimDeviceBaysBulkPartialUpdateOK {
	return &DcimDeviceBaysBulkPartialUpdateOK{}
}

/* DcimDeviceBaysBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBaysBulkPartialUpdateOK dcim device bays bulk partial update o k
*/
type DcimDeviceBaysBulkPartialUpdateOK struct {
	Payload *models.DeviceBay
}

func (o *DcimDeviceBaysBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bays/][%d] dcimDeviceBaysBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceBaysBulkPartialUpdateOK) GetPayload() *models.DeviceBay {
	return o.Payload
}

func (o *DcimDeviceBaysBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
