package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerFeedsUpdateReader is a Reader for the DcimPowerFeedsUpdate structure.
type DcimPowerFeedsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerFeedsUpdateOK creates a DcimPowerFeedsUpdateOK with default headers values
func NewDcimPowerFeedsUpdateOK() *DcimPowerFeedsUpdateOK {
	return &DcimPowerFeedsUpdateOK{}
}

/* DcimPowerFeedsUpdateOK describes a response with status code 200, with default header values.

DcimPowerFeedsUpdateOK dcim power feeds update o k
*/
type DcimPowerFeedsUpdateOK struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/{id}/][%d] dcimPowerFeedsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerFeedsUpdateOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
