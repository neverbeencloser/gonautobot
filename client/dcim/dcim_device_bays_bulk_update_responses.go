package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceBaysBulkUpdateReader is a Reader for the DcimDeviceBaysBulkUpdate structure.
type DcimDeviceBaysBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBaysBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBaysBulkUpdateOK creates a DcimDeviceBaysBulkUpdateOK with default headers values
func NewDcimDeviceBaysBulkUpdateOK() *DcimDeviceBaysBulkUpdateOK {
	return &DcimDeviceBaysBulkUpdateOK{}
}

/* DcimDeviceBaysBulkUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBaysBulkUpdateOK dcim device bays bulk update o k
*/
type DcimDeviceBaysBulkUpdateOK struct {
	Payload *models.DeviceBay
}

func (o *DcimDeviceBaysBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bays/][%d] dcimDeviceBaysBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceBaysBulkUpdateOK) GetPayload() *models.DeviceBay {
	return o.Payload
}

func (o *DcimDeviceBaysBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
