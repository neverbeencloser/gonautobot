package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerFeedsCreateReader is a Reader for the DcimPowerFeedsCreate structure.
type DcimPowerFeedsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerFeedsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerFeedsCreateCreated creates a DcimPowerFeedsCreateCreated with default headers values
func NewDcimPowerFeedsCreateCreated() *DcimPowerFeedsCreateCreated {
	return &DcimPowerFeedsCreateCreated{}
}

/* DcimPowerFeedsCreateCreated describes a response with status code 201, with default header values.

DcimPowerFeedsCreateCreated dcim power feeds create created
*/
type DcimPowerFeedsCreateCreated struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-feeds/][%d] dcimPowerFeedsCreateCreated  %+v", 201, o.Payload)
}
func (o *DcimPowerFeedsCreateCreated) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
