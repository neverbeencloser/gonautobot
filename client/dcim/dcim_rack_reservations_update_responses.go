package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackReservationsUpdateReader is a Reader for the DcimRackReservationsUpdate structure.
type DcimRackReservationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackReservationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackReservationsUpdateOK creates a DcimRackReservationsUpdateOK with default headers values
func NewDcimRackReservationsUpdateOK() *DcimRackReservationsUpdateOK {
	return &DcimRackReservationsUpdateOK{}
}

/* DcimRackReservationsUpdateOK describes a response with status code 200, with default header values.

DcimRackReservationsUpdateOK dcim rack reservations update o k
*/
type DcimRackReservationsUpdateOK struct {
	Payload *models.RackReservation
}

func (o *DcimRackReservationsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-reservations/{id}/][%d] dcimRackReservationsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRackReservationsUpdateOK) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
