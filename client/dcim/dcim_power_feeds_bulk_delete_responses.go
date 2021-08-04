package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerFeedsBulkDeleteReader is a Reader for the DcimPowerFeedsBulkDelete structure.
type DcimPowerFeedsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerFeedsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerFeedsBulkDeleteNoContent creates a DcimPowerFeedsBulkDeleteNoContent with default headers values
func NewDcimPowerFeedsBulkDeleteNoContent() *DcimPowerFeedsBulkDeleteNoContent {
	return &DcimPowerFeedsBulkDeleteNoContent{}
}

/* DcimPowerFeedsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerFeedsBulkDeleteNoContent dcim power feeds bulk delete no content
*/
type DcimPowerFeedsBulkDeleteNoContent struct {
}

func (o *DcimPowerFeedsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-feeds/][%d] dcimPowerFeedsBulkDeleteNoContent ", 204)
}

func (o *DcimPowerFeedsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
