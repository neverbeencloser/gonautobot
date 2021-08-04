package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerOutletTemplatesReadReader is a Reader for the DcimPowerOutletTemplatesRead structure.
type DcimPowerOutletTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesReadOK creates a DcimPowerOutletTemplatesReadOK with default headers values
func NewDcimPowerOutletTemplatesReadOK() *DcimPowerOutletTemplatesReadOK {
	return &DcimPowerOutletTemplatesReadOK{}
}

/* DcimPowerOutletTemplatesReadOK describes a response with status code 200, with default header values.

DcimPowerOutletTemplatesReadOK dcim power outlet templates read o k
*/
type DcimPowerOutletTemplatesReadOK struct {
	Payload *models.PowerOutletTemplate
}

func (o *DcimPowerOutletTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesReadOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletTemplatesReadOK) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
