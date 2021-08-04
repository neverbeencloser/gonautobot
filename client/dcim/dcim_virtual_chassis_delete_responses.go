package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimVirtualChassisDeleteReader is a Reader for the DcimVirtualChassisDelete structure.
type DcimVirtualChassisDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimVirtualChassisDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimVirtualChassisDeleteNoContent creates a DcimVirtualChassisDeleteNoContent with default headers values
func NewDcimVirtualChassisDeleteNoContent() *DcimVirtualChassisDeleteNoContent {
	return &DcimVirtualChassisDeleteNoContent{}
}

/* DcimVirtualChassisDeleteNoContent describes a response with status code 204, with default header values.

DcimVirtualChassisDeleteNoContent dcim virtual chassis delete no content
*/
type DcimVirtualChassisDeleteNoContent struct {
}

func (o *DcimVirtualChassisDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisDeleteNoContent ", 204)
}

func (o *DcimVirtualChassisDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
