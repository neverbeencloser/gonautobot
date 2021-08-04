package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceBayTemplatesBulkDeleteReader is a Reader for the DcimDeviceBayTemplatesBulkDelete structure.
type DcimDeviceBayTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceBayTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesBulkDeleteNoContent creates a DcimDeviceBayTemplatesBulkDeleteNoContent with default headers values
func NewDcimDeviceBayTemplatesBulkDeleteNoContent() *DcimDeviceBayTemplatesBulkDeleteNoContent {
	return &DcimDeviceBayTemplatesBulkDeleteNoContent{}
}

/* DcimDeviceBayTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceBayTemplatesBulkDeleteNoContent dcim device bay templates bulk delete no content
*/
type DcimDeviceBayTemplatesBulkDeleteNoContent struct {
}

func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkDeleteNoContent ", 204)
}

func (o *DcimDeviceBayTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
