package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceBayTemplatesUpdateReader is a Reader for the DcimDeviceBayTemplatesUpdate structure.
type DcimDeviceBayTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesUpdateOK creates a DcimDeviceBayTemplatesUpdateOK with default headers values
func NewDcimDeviceBayTemplatesUpdateOK() *DcimDeviceBayTemplatesUpdateOK {
	return &DcimDeviceBayTemplatesUpdateOK{}
}

/* DcimDeviceBayTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesUpdateOK dcim device bay templates update o k
*/
type DcimDeviceBayTemplatesUpdateOK struct {
	Payload *models.DeviceBayTemplate
}

func (o *DcimDeviceBayTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceBayTemplatesUpdateOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
