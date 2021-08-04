package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRackGroupsBulkDeleteReader is a Reader for the DcimRackGroupsBulkDelete structure.
type DcimRackGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRackGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackGroupsBulkDeleteNoContent creates a DcimRackGroupsBulkDeleteNoContent with default headers values
func NewDcimRackGroupsBulkDeleteNoContent() *DcimRackGroupsBulkDeleteNoContent {
	return &DcimRackGroupsBulkDeleteNoContent{}
}

/* DcimRackGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRackGroupsBulkDeleteNoContent dcim rack groups bulk delete no content
*/
type DcimRackGroupsBulkDeleteNoContent struct {
}

func (o *DcimRackGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-groups/][%d] dcimRackGroupsBulkDeleteNoContent ", 204)
}

func (o *DcimRackGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
