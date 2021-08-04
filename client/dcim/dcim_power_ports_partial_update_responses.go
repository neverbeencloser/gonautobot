package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPortsPartialUpdateReader is a Reader for the DcimPowerPortsPartialUpdate structure.
type DcimPowerPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortsPartialUpdateOK creates a DcimPowerPortsPartialUpdateOK with default headers values
func NewDcimPowerPortsPartialUpdateOK() *DcimPowerPortsPartialUpdateOK {
	return &DcimPowerPortsPartialUpdateOK{}
}

/* DcimPowerPortsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortsPartialUpdateOK dcim power ports partial update o k
*/
type DcimPowerPortsPartialUpdateOK struct {
	Payload *models.PowerPort
}

func (o *DcimPowerPortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-ports/{id}/][%d] dcimPowerPortsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPortsPartialUpdateOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
