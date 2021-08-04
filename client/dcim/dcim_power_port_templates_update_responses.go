package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPortTemplatesUpdateReader is a Reader for the DcimPowerPortTemplatesUpdate structure.
type DcimPowerPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesUpdateOK creates a DcimPowerPortTemplatesUpdateOK with default headers values
func NewDcimPowerPortTemplatesUpdateOK() *DcimPowerPortTemplatesUpdateOK {
	return &DcimPowerPortTemplatesUpdateOK{}
}

/* DcimPowerPortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesUpdateOK dcim power port templates update o k
*/
type DcimPowerPortTemplatesUpdateOK struct {
	Payload *models.PowerPortTemplate
}

func (o *DcimPowerPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPortTemplatesUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
