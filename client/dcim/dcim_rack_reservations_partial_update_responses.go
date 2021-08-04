package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackReservationsPartialUpdateReader is a Reader for the DcimRackReservationsPartialUpdate structure.
type DcimRackReservationsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackReservationsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackReservationsPartialUpdateOK creates a DcimRackReservationsPartialUpdateOK with default headers values
func NewDcimRackReservationsPartialUpdateOK() *DcimRackReservationsPartialUpdateOK {
	return &DcimRackReservationsPartialUpdateOK{}
}

/* DcimRackReservationsPartialUpdateOK describes a response with status code 200, with default header values.

DcimRackReservationsPartialUpdateOK dcim rack reservations partial update o k
*/
type DcimRackReservationsPartialUpdateOK struct {
	Payload *models.RackReservation
}

func (o *DcimRackReservationsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-reservations/{id}/][%d] dcimRackReservationsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRackReservationsPartialUpdateOK) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
