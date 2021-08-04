package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceBayTemplatesDeleteReader is a Reader for the DcimDeviceBayTemplatesDelete structure.
type DcimDeviceBayTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceBayTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesDeleteNoContent creates a DcimDeviceBayTemplatesDeleteNoContent with default headers values
func NewDcimDeviceBayTemplatesDeleteNoContent() *DcimDeviceBayTemplatesDeleteNoContent {
	return &DcimDeviceBayTemplatesDeleteNoContent{}
}

/* DcimDeviceBayTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceBayTemplatesDeleteNoContent dcim device bay templates delete no content
*/
type DcimDeviceBayTemplatesDeleteNoContent struct {
}

func (o *DcimDeviceBayTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesDeleteNoContent ", 204)
}

func (o *DcimDeviceBayTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
