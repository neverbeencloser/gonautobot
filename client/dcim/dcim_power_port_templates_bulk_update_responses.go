package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPortTemplatesBulkUpdateReader is a Reader for the DcimPowerPortTemplatesBulkUpdate structure.
type DcimPowerPortTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesBulkUpdateOK creates a DcimPowerPortTemplatesBulkUpdateOK with default headers values
func NewDcimPowerPortTemplatesBulkUpdateOK() *DcimPowerPortTemplatesBulkUpdateOK {
	return &DcimPowerPortTemplatesBulkUpdateOK{}
}

/* DcimPowerPortTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesBulkUpdateOK dcim power port templates bulk update o k
*/
type DcimPowerPortTemplatesBulkUpdateOK struct {
	Payload *models.PowerPortTemplate
}

func (o *DcimPowerPortTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/][%d] dcimPowerPortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPortTemplatesBulkUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
