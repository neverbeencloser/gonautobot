package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRackReservationsDeleteReader is a Reader for the DcimRackReservationsDelete structure.
type DcimRackReservationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRackReservationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackReservationsDeleteNoContent creates a DcimRackReservationsDeleteNoContent with default headers values
func NewDcimRackReservationsDeleteNoContent() *DcimRackReservationsDeleteNoContent {
	return &DcimRackReservationsDeleteNoContent{}
}

/* DcimRackReservationsDeleteNoContent describes a response with status code 204, with default header values.

DcimRackReservationsDeleteNoContent dcim rack reservations delete no content
*/
type DcimRackReservationsDeleteNoContent struct {
}

func (o *DcimRackReservationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-reservations/{id}/][%d] dcimRackReservationsDeleteNoContent ", 204)
}

func (o *DcimRackReservationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
