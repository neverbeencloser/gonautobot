package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceBaysBulkDeleteReader is a Reader for the DcimDeviceBaysBulkDelete structure.
type DcimDeviceBaysBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceBaysBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBaysBulkDeleteNoContent creates a DcimDeviceBaysBulkDeleteNoContent with default headers values
func NewDcimDeviceBaysBulkDeleteNoContent() *DcimDeviceBaysBulkDeleteNoContent {
	return &DcimDeviceBaysBulkDeleteNoContent{}
}

/* DcimDeviceBaysBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceBaysBulkDeleteNoContent dcim device bays bulk delete no content
*/
type DcimDeviceBaysBulkDeleteNoContent struct {
}

func (o *DcimDeviceBaysBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bays/][%d] dcimDeviceBaysBulkDeleteNoContent ", 204)
}

func (o *DcimDeviceBaysBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
