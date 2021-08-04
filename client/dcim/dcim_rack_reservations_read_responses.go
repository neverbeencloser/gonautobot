package dcim

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// DcimRackReservationsReadReader is a Reader for the DcimRackReservationsRead structure.
type DcimRackReservationsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackReservationsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackReservationsReadOK creates a DcimRackReservationsReadOK with default headers values
func NewDcimRackReservationsReadOK() *DcimRackReservationsReadOK {
	return &DcimRackReservationsReadOK{}
}

/* DcimRackReservationsReadOK describes a response with status code 200, with default header values.

DcimRackReservationsReadOK dcim rack reservations read o k
*/
type DcimRackReservationsReadOK struct {
	Payload *models.RackReservation
}

func (o *DcimRackReservationsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-reservations/{id}/][%d] dcimRackReservationsReadOK  %+v", 200, o.Payload)
}
func (o *DcimRackReservationsReadOK) GetPayload() *models.RackReservation {
	return o.Payload
}

func (o *DcimRackReservationsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
