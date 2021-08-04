package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRacksBulkDeleteReader is a Reader for the DcimRacksBulkDelete structure.
type DcimRacksBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRacksBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRacksBulkDeleteNoContent creates a DcimRacksBulkDeleteNoContent with default headers values
func NewDcimRacksBulkDeleteNoContent() *DcimRacksBulkDeleteNoContent {
	return &DcimRacksBulkDeleteNoContent{}
}

/* DcimRacksBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRacksBulkDeleteNoContent dcim racks bulk delete no content
*/
type DcimRacksBulkDeleteNoContent struct {
}

func (o *DcimRacksBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/racks/][%d] dcimRacksBulkDeleteNoContent ", 204)
}

func (o *DcimRacksBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
