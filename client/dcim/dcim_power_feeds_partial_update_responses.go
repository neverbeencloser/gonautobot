package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerFeedsPartialUpdateReader is a Reader for the DcimPowerFeedsPartialUpdate structure.
type DcimPowerFeedsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerFeedsPartialUpdateOK creates a DcimPowerFeedsPartialUpdateOK with default headers values
func NewDcimPowerFeedsPartialUpdateOK() *DcimPowerFeedsPartialUpdateOK {
	return &DcimPowerFeedsPartialUpdateOK{}
}

/* DcimPowerFeedsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerFeedsPartialUpdateOK dcim power feeds partial update o k
*/
type DcimPowerFeedsPartialUpdateOK struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-feeds/{id}/][%d] dcimPowerFeedsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerFeedsPartialUpdateOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
