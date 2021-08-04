package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerFeedsReadReader is a Reader for the DcimPowerFeedsRead structure.
type DcimPowerFeedsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerFeedsReadOK creates a DcimPowerFeedsReadOK with default headers values
func NewDcimPowerFeedsReadOK() *DcimPowerFeedsReadOK {
	return &DcimPowerFeedsReadOK{}
}

/* DcimPowerFeedsReadOK describes a response with status code 200, with default header values.

DcimPowerFeedsReadOK dcim power feeds read o k
*/
type DcimPowerFeedsReadOK struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/][%d] dcimPowerFeedsReadOK  %+v", 200, o.Payload)
}
func (o *DcimPowerFeedsReadOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
