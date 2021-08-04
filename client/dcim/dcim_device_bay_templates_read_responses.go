package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceBayTemplatesReadReader is a Reader for the DcimDeviceBayTemplatesRead structure.
type DcimDeviceBayTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesReadOK creates a DcimDeviceBayTemplatesReadOK with default headers values
func NewDcimDeviceBayTemplatesReadOK() *DcimDeviceBayTemplatesReadOK {
	return &DcimDeviceBayTemplatesReadOK{}
}

/* DcimDeviceBayTemplatesReadOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesReadOK dcim device bay templates read o k
*/
type DcimDeviceBayTemplatesReadOK struct {
	Payload *models.DeviceBayTemplate
}

func (o *DcimDeviceBayTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesReadOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceBayTemplatesReadOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
