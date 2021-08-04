package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerOutletsTraceReader is a Reader for the DcimPowerOutletsTrace structure.
type DcimPowerOutletsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletsTraceOK creates a DcimPowerOutletsTraceOK with default headers values
func NewDcimPowerOutletsTraceOK() *DcimPowerOutletsTraceOK {
	return &DcimPowerOutletsTraceOK{}
}

/* DcimPowerOutletsTraceOK describes a response with status code 200, with default header values.

DcimPowerOutletsTraceOK dcim power outlets trace o k
*/
type DcimPowerOutletsTraceOK struct {
	Payload *models.PowerOutlet
}

func (o *DcimPowerOutletsTraceOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/trace/][%d] dcimPowerOutletsTraceOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletsTraceOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
