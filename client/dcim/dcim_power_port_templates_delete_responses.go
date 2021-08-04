package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerPortTemplatesDeleteReader is a Reader for the DcimPowerPortTemplatesDelete structure.
type DcimPowerPortTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPortTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesDeleteNoContent creates a DcimPowerPortTemplatesDeleteNoContent with default headers values
func NewDcimPowerPortTemplatesDeleteNoContent() *DcimPowerPortTemplatesDeleteNoContent {
	return &DcimPowerPortTemplatesDeleteNoContent{}
}

/* DcimPowerPortTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPortTemplatesDeleteNoContent dcim power port templates delete no content
*/
type DcimPowerPortTemplatesDeleteNoContent struct {
}

func (o *DcimPowerPortTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimPowerPortTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
