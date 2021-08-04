package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerPanelsPartialUpdateReader is a Reader for the DcimPowerPanelsPartialUpdate structure.
type DcimPowerPanelsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPanelsPartialUpdateOK creates a DcimPowerPanelsPartialUpdateOK with default headers values
func NewDcimPowerPanelsPartialUpdateOK() *DcimPowerPanelsPartialUpdateOK {
	return &DcimPowerPanelsPartialUpdateOK{}
}

/* DcimPowerPanelsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPanelsPartialUpdateOK dcim power panels partial update o k
*/
type DcimPowerPanelsPartialUpdateOK struct {
	Payload *models.PowerPanel
}

func (o *DcimPowerPanelsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/{id}/][%d] dcimPowerPanelsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerPanelsPartialUpdateOK) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
