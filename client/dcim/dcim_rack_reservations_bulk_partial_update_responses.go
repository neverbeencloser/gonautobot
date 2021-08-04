package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackReservationsBulkPartialUpdateReader is a Reader for the DcimRackReservationsBulkPartialUpdate structure.
type DcimRackReservationsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackReservationsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackReservationsBulkPartialUpdateOK creates a DcimRackReservationsBulkPartialUpdateOK with default headers values
func NewDcimRackReservationsBulkPartialUpdateOK() *DcimRackReservationsBulkPartialUpdateOK {
	return &DcimRackReservationsBulkPartialUpdateOK{}
}

/* DcimRackReservationsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRackReservationsBulkPartialUpdateOK dcim rack reservations bulk partial update o k
*/
type DcimRackReservationsBulkPartialUpdateOK struct {
	Payload *models.RackReservation
}

func (o *DcimRackReservationsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-reservations/][%d] dcimRackReservationsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRackReservationsBulkPartialUpdateOK) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
