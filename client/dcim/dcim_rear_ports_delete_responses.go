package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRearPortsDeleteReader is a Reader for the DcimRearPortsDelete structure.
type DcimRearPortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRearPortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortsDeleteNoContent creates a DcimRearPortsDeleteNoContent with default headers values
func NewDcimRearPortsDeleteNoContent() *DcimRearPortsDeleteNoContent {
	return &DcimRearPortsDeleteNoContent{}
}

/* DcimRearPortsDeleteNoContent describes a response with status code 204, with default header values.

DcimRearPortsDeleteNoContent dcim rear ports delete no content
*/
type DcimRearPortsDeleteNoContent struct {
}

func (o *DcimRearPortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rear-ports/{id}/][%d] dcimRearPortsDeleteNoContent ", 204)
}

func (o *DcimRearPortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
