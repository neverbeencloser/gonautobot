package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimPowerOutletsBulkPartialUpdateReader is a Reader for the DcimPowerOutletsBulkPartialUpdate structure.
type DcimPowerOutletsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletsBulkPartialUpdateOK creates a DcimPowerOutletsBulkPartialUpdateOK with default headers values
func NewDcimPowerOutletsBulkPartialUpdateOK() *DcimPowerOutletsBulkPartialUpdateOK {
	return &DcimPowerOutletsBulkPartialUpdateOK{}
}

/* DcimPowerOutletsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletsBulkPartialUpdateOK dcim power outlets bulk partial update o k
*/
type DcimPowerOutletsBulkPartialUpdateOK struct {
	Payload *models.PowerOutlet
}

func (o *DcimPowerOutletsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlets/][%d] dcimPowerOutletsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerOutletsBulkPartialUpdateOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
