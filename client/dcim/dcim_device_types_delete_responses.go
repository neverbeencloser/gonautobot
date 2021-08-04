package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceTypesDeleteReader is a Reader for the DcimDeviceTypesDelete structure.
type DcimDeviceTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceTypesDeleteNoContent creates a DcimDeviceTypesDeleteNoContent with default headers values
func NewDcimDeviceTypesDeleteNoContent() *DcimDeviceTypesDeleteNoContent {
	return &DcimDeviceTypesDeleteNoContent{}
}

/* DcimDeviceTypesDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceTypesDeleteNoContent dcim device types delete no content
*/
type DcimDeviceTypesDeleteNoContent struct {
}

func (o *DcimDeviceTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/{id}/][%d] dcimDeviceTypesDeleteNoContent ", 204)
}

func (o *DcimDeviceTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
