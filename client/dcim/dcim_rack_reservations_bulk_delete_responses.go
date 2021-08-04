package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRackReservationsBulkDeleteReader is a Reader for the DcimRackReservationsBulkDelete structure.
type DcimRackReservationsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRackReservationsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackReservationsBulkDeleteNoContent creates a DcimRackReservationsBulkDeleteNoContent with default headers values
func NewDcimRackReservationsBulkDeleteNoContent() *DcimRackReservationsBulkDeleteNoContent {
	return &DcimRackReservationsBulkDeleteNoContent{}
}

/* DcimRackReservationsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRackReservationsBulkDeleteNoContent dcim rack reservations bulk delete no content
*/
type DcimRackReservationsBulkDeleteNoContent struct {
}

func (o *DcimRackReservationsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-reservations/][%d] dcimRackReservationsBulkDeleteNoContent ", 204)
}

func (o *DcimRackReservationsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
