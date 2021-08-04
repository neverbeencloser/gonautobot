package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerOutletTemplatesBulkPartialUpdateReader is a Reader for the DcimPowerOutletTemplatesBulkPartialUpdate structure.
type DcimPowerOutletTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesBulkPartialUpdateOK creates a DcimPowerOutletTemplatesBulkPartialUpdateOK with default headers values
func NewDcimPowerOutletTemplatesBulkPartialUpdateOK() *DcimPowerOutletTemplatesBulkPartialUpdateOK {
	return &DcimPowerOutletTemplatesBulkPartialUpdateOK{}
}

/* DcimPowerOutletTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletTemplatesBulkPartialUpdateOK dcim power outlet templates bulk partial update o k
*/
type DcimPowerOutletTemplatesBulkPartialUpdateOK struct {
	Payload *models.PowerOutletTemplate
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
