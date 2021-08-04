package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPortTemplatesPartialUpdateReader is a Reader for the DcimPowerPortTemplatesPartialUpdate structure.
type DcimPowerPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesPartialUpdateOK creates a DcimPowerPortTemplatesPartialUpdateOK with default headers values
func NewDcimPowerPortTemplatesPartialUpdateOK() *DcimPowerPortTemplatesPartialUpdateOK {
	return &DcimPowerPortTemplatesPartialUpdateOK{}
}

/* DcimPowerPortTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesPartialUpdateOK dcim power port templates partial update o k
*/
type DcimPowerPortTemplatesPartialUpdateOK struct {
	Payload *models.PowerPortTemplate
}

func (o *DcimPowerPortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPortTemplatesPartialUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
