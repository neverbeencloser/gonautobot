package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDevicesDeleteReader is a Reader for the DcimDevicesDelete structure.
type DcimDevicesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDevicesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDevicesDeleteNoContent creates a DcimDevicesDeleteNoContent with default headers values
func NewDcimDevicesDeleteNoContent() *DcimDevicesDeleteNoContent {
	return &DcimDevicesDeleteNoContent{}
}

/* DcimDevicesDeleteNoContent describes a response with status code 204, with default header values.

DcimDevicesDeleteNoContent dcim devices delete no content
*/
type DcimDevicesDeleteNoContent struct {
}

func (o *DcimDevicesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/devices/{id}/][%d] dcimDevicesDeleteNoContent ", 204)
}

func (o *DcimDevicesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
