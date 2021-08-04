package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPanelsCreateReader is a Reader for the DcimPowerPanelsCreate structure.
type DcimPowerPanelsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerPanelsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPanelsCreateCreated creates a DcimPowerPanelsCreateCreated with default headers values
func NewDcimPowerPanelsCreateCreated() *DcimPowerPanelsCreateCreated {
	return &DcimPowerPanelsCreateCreated{}
}

/* DcimPowerPanelsCreateCreated describes a response with status code 201, with default header values.

DcimPowerPanelsCreateCreated dcim power panels create created
*/
type DcimPowerPanelsCreateCreated struct {
	Payload *models.PowerPanel
}

func (o *DcimPowerPanelsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-panels/][%d] dcimPowerPanelsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimPowerPanelsCreateCreated) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
