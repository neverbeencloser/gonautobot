package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRackRolesBulkDeleteReader is a Reader for the DcimRackRolesBulkDelete structure.
type DcimRackRolesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRackRolesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackRolesBulkDeleteNoContent creates a DcimRackRolesBulkDeleteNoContent with default headers values
func NewDcimRackRolesBulkDeleteNoContent() *DcimRackRolesBulkDeleteNoContent {
	return &DcimRackRolesBulkDeleteNoContent{}
}

/* DcimRackRolesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRackRolesBulkDeleteNoContent dcim rack roles bulk delete no content
*/
type DcimRackRolesBulkDeleteNoContent struct {
}

func (o *DcimRackRolesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-roles/][%d] dcimRackRolesBulkDeleteNoContent ", 204)
}

func (o *DcimRackRolesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
