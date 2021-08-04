package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerOutletsBulkUpdateReader is a Reader for the DcimPowerOutletsBulkUpdate structure.
type DcimPowerOutletsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletsBulkUpdateOK creates a DcimPowerOutletsBulkUpdateOK with default headers values
func NewDcimPowerOutletsBulkUpdateOK() *DcimPowerOutletsBulkUpdateOK {
	return &DcimPowerOutletsBulkUpdateOK{}
}

/* DcimPowerOutletsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletsBulkUpdateOK dcim power outlets bulk update o k
*/
type DcimPowerOutletsBulkUpdateOK struct {
	Payload *models.PowerOutlet
}

func (o *DcimPowerOutletsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/][%d] dcimPowerOutletsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletsBulkUpdateOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
