package dcim

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimVirtualChassisBulkDeleteReader is a Reader for the DcimVirtualChassisBulkDelete structure.
type DcimVirtualChassisBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimVirtualChassisBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimVirtualChassisBulkDeleteNoContent creates a DcimVirtualChassisBulkDeleteNoContent with default headers values
func NewDcimVirtualChassisBulkDeleteNoContent() *DcimVirtualChassisBulkDeleteNoContent {
	return &DcimVirtualChassisBulkDeleteNoContent{}
}

/* DcimVirtualChassisBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimVirtualChassisBulkDeleteNoContent dcim virtual chassis bulk delete no content
*/
type DcimVirtualChassisBulkDeleteNoContent struct {
}

func (o *DcimVirtualChassisBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/virtual-chassis/][%d] dcimVirtualChassisBulkDeleteNoContent ", 204)
}

func (o *DcimVirtualChassisBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
