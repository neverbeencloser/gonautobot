package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamRolesBulkDeleteReader is a Reader for the IpamRolesBulkDelete structure.
type IpamRolesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRolesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRolesBulkDeleteNoContent creates a IpamRolesBulkDeleteNoContent with default headers values
func NewIpamRolesBulkDeleteNoContent() *IpamRolesBulkDeleteNoContent {
	return &IpamRolesBulkDeleteNoContent{}
}

/* IpamRolesBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamRolesBulkDeleteNoContent ipam roles bulk delete no content
*/
type IpamRolesBulkDeleteNoContent struct {
}

func (o *IpamRolesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/roles/][%d] ipamRolesBulkDeleteNoContent ", 204)
}

func (o *IpamRolesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
