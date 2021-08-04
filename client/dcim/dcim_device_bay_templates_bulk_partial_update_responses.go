package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimDeviceBayTemplatesBulkPartialUpdateReader is a Reader for the DcimDeviceBayTemplatesBulkPartialUpdate structure.
type DcimDeviceBayTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesBulkPartialUpdateOK creates a DcimDeviceBayTemplatesBulkPartialUpdateOK with default headers values
func NewDcimDeviceBayTemplatesBulkPartialUpdateOK() *DcimDeviceBayTemplatesBulkPartialUpdateOK {
	return &DcimDeviceBayTemplatesBulkPartialUpdateOK{}
}

/* DcimDeviceBayTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesBulkPartialUpdateOK dcim device bay templates bulk partial update o k
*/
type DcimDeviceBayTemplatesBulkPartialUpdateOK struct {
	Payload *models.DeviceBayTemplate
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
