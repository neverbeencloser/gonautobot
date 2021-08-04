package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPortsCreateReader is a Reader for the DcimPowerPortsCreate structure.
type DcimPowerPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortsCreateCreated creates a DcimPowerPortsCreateCreated with default headers values
func NewDcimPowerPortsCreateCreated() *DcimPowerPortsCreateCreated {
	return &DcimPowerPortsCreateCreated{}
}

/* DcimPowerPortsCreateCreated describes a response with status code 201, with default header values.

DcimPowerPortsCreateCreated dcim power ports create created
*/
type DcimPowerPortsCreateCreated struct {
	Payload *models.PowerPort
}

func (o *DcimPowerPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-ports/][%d] dcimPowerPortsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimPowerPortsCreateCreated) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
