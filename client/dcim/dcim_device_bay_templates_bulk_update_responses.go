package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceBayTemplatesBulkUpdateReader is a Reader for the DcimDeviceBayTemplatesBulkUpdate structure.
type DcimDeviceBayTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesBulkUpdateOK creates a DcimDeviceBayTemplatesBulkUpdateOK with default headers values
func NewDcimDeviceBayTemplatesBulkUpdateOK() *DcimDeviceBayTemplatesBulkUpdateOK {
	return &DcimDeviceBayTemplatesBulkUpdateOK{}
}

/* DcimDeviceBayTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesBulkUpdateOK dcim device bay templates bulk update o k
*/
type DcimDeviceBayTemplatesBulkUpdateOK struct {
	Payload *models.DeviceBayTemplate
}

func (o *DcimDeviceBayTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceBayTemplatesBulkUpdateOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
