package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerPanelsDeleteReader is a Reader for the DcimPowerPanelsDelete structure.
type DcimPowerPanelsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPanelsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPanelsDeleteNoContent creates a DcimPowerPanelsDeleteNoContent with default headers values
func NewDcimPowerPanelsDeleteNoContent() *DcimPowerPanelsDeleteNoContent {
	return &DcimPowerPanelsDeleteNoContent{}
}

/* DcimPowerPanelsDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPanelsDeleteNoContent dcim power panels delete no content
*/
type DcimPowerPanelsDeleteNoContent struct {
}

func (o *DcimPowerPanelsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-panels/{id}/][%d] dcimPowerPanelsDeleteNoContent ", 204)
}

func (o *DcimPowerPanelsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
