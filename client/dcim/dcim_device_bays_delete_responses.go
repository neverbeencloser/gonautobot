package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceBaysDeleteReader is a Reader for the DcimDeviceBaysDelete structure.
type DcimDeviceBaysDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceBaysDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBaysDeleteNoContent creates a DcimDeviceBaysDeleteNoContent with default headers values
func NewDcimDeviceBaysDeleteNoContent() *DcimDeviceBaysDeleteNoContent {
	return &DcimDeviceBaysDeleteNoContent{}
}

/* DcimDeviceBaysDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceBaysDeleteNoContent dcim device bays delete no content
*/
type DcimDeviceBaysDeleteNoContent struct {
}

func (o *DcimDeviceBaysDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bays/{id}/][%d] dcimDeviceBaysDeleteNoContent ", 204)
}

func (o *DcimDeviceBaysDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
