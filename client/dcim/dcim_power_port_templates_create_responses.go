package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPortTemplatesCreateReader is a Reader for the DcimPowerPortTemplatesCreate structure.
type DcimPowerPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesCreateCreated creates a DcimPowerPortTemplatesCreateCreated with default headers values
func NewDcimPowerPortTemplatesCreateCreated() *DcimPowerPortTemplatesCreateCreated {
	return &DcimPowerPortTemplatesCreateCreated{}
}

/* DcimPowerPortTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimPowerPortTemplatesCreateCreated dcim power port templates create created
*/
type DcimPowerPortTemplatesCreateCreated struct {
	Payload *models.PowerPortTemplate
}

func (o *DcimPowerPortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-port-templates/][%d] dcimPowerPortTemplatesCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimPowerPortTemplatesCreateCreated) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
