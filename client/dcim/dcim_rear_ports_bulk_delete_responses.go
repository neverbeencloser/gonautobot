package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRearPortsBulkDeleteReader is a Reader for the DcimRearPortsBulkDelete structure.
type DcimRearPortsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRearPortsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortsBulkDeleteNoContent creates a DcimRearPortsBulkDeleteNoContent with default headers values
func NewDcimRearPortsBulkDeleteNoContent() *DcimRearPortsBulkDeleteNoContent {
	return &DcimRearPortsBulkDeleteNoContent{}
}

/* DcimRearPortsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRearPortsBulkDeleteNoContent dcim rear ports bulk delete no content
*/
type DcimRearPortsBulkDeleteNoContent struct {
}

func (o *DcimRearPortsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rear-ports/][%d] dcimRearPortsBulkDeleteNoContent ", 204)
}

func (o *DcimRearPortsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
