package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerFeedsTraceReader is a Reader for the DcimPowerFeedsTrace structure.
type DcimPowerFeedsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerFeedsTraceOK creates a DcimPowerFeedsTraceOK with default headers values
func NewDcimPowerFeedsTraceOK() *DcimPowerFeedsTraceOK {
	return &DcimPowerFeedsTraceOK{}
}

/* DcimPowerFeedsTraceOK describes a response with status code 200, with default header values.

DcimPowerFeedsTraceOK dcim power feeds trace o k
*/
type DcimPowerFeedsTraceOK struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsTraceOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-feeds/{id}/trace/][%d] dcimPowerFeedsTraceOK  %+v", 200, o.Payload)
}
func (o *DcimPowerFeedsTraceOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
