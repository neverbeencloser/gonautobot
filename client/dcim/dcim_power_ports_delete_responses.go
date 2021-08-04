package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerPortsDeleteReader is a Reader for the DcimPowerPortsDelete structure.
type DcimPowerPortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortsDeleteNoContent creates a DcimPowerPortsDeleteNoContent with default headers values
func NewDcimPowerPortsDeleteNoContent() *DcimPowerPortsDeleteNoContent {
	return &DcimPowerPortsDeleteNoContent{}
}

/* DcimPowerPortsDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPortsDeleteNoContent dcim power ports delete no content
*/
type DcimPowerPortsDeleteNoContent struct {
}

func (o *DcimPowerPortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-ports/{id}/][%d] dcimPowerPortsDeleteNoContent ", 204)
}

func (o *DcimPowerPortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
