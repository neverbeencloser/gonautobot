package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVlansBulkDeleteReader is a Reader for the IpamVlansBulkDelete structure.
type IpamVlansBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVlansBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlansBulkDeleteNoContent creates a IpamVlansBulkDeleteNoContent with default headers values
func NewIpamVlansBulkDeleteNoContent() *IpamVlansBulkDeleteNoContent {
	return &IpamVlansBulkDeleteNoContent{}
}

/* IpamVlansBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamVlansBulkDeleteNoContent ipam vlans bulk delete no content
*/
type IpamVlansBulkDeleteNoContent struct {
}

func (o *IpamVlansBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlans/][%d] ipamVlansBulkDeleteNoContent ", 204)
}

func (o *IpamVlansBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
