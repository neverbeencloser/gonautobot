package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerOutletsCreateReader is a Reader for the DcimPowerOutletsCreate structure.
type DcimPowerOutletsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerOutletsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletsCreateCreated creates a DcimPowerOutletsCreateCreated with default headers values
func NewDcimPowerOutletsCreateCreated() *DcimPowerOutletsCreateCreated {
	return &DcimPowerOutletsCreateCreated{}
}

/* DcimPowerOutletsCreateCreated describes a response with status code 201, with default header values.

DcimPowerOutletsCreateCreated dcim power outlets create created
*/
type DcimPowerOutletsCreateCreated struct {
	Payload *models.PowerOutlet
}

func (o *DcimPowerOutletsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-outlets/][%d] dcimPowerOutletsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimPowerOutletsCreateCreated) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
