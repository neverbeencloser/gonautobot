package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerOutletTemplatesBulkUpdateReader is a Reader for the DcimPowerOutletTemplatesBulkUpdate structure.
type DcimPowerOutletTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesBulkUpdateOK creates a DcimPowerOutletTemplatesBulkUpdateOK with default headers values
func NewDcimPowerOutletTemplatesBulkUpdateOK() *DcimPowerOutletTemplatesBulkUpdateOK {
	return &DcimPowerOutletTemplatesBulkUpdateOK{}
}

/* DcimPowerOutletTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletTemplatesBulkUpdateOK dcim power outlet templates bulk update o k
*/
type DcimPowerOutletTemplatesBulkUpdateOK struct {
	Payload *models.PowerOutletTemplate
}

func (o *DcimPowerOutletTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletTemplatesBulkUpdateOK) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
