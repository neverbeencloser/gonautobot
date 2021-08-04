package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerFeedsBulkUpdateReader is a Reader for the DcimPowerFeedsBulkUpdate structure.
type DcimPowerFeedsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerFeedsBulkUpdateOK creates a DcimPowerFeedsBulkUpdateOK with default headers values
func NewDcimPowerFeedsBulkUpdateOK() *DcimPowerFeedsBulkUpdateOK {
	return &DcimPowerFeedsBulkUpdateOK{}
}

/* DcimPowerFeedsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerFeedsBulkUpdateOK dcim power feeds bulk update o k
*/
type DcimPowerFeedsBulkUpdateOK struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcimPowerFeedsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerFeedsBulkUpdateOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
