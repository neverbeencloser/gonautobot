package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVlanGroupsBulkDeleteReader is a Reader for the IpamVlanGroupsBulkDelete structure.
type IpamVlanGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVlanGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlanGroupsBulkDeleteNoContent creates a IpamVlanGroupsBulkDeleteNoContent with default headers values
func NewIpamVlanGroupsBulkDeleteNoContent() *IpamVlanGroupsBulkDeleteNoContent {
	return &IpamVlanGroupsBulkDeleteNoContent{}
}

/* IpamVlanGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamVlanGroupsBulkDeleteNoContent ipam vlan groups bulk delete no content
*/
type IpamVlanGroupsBulkDeleteNoContent struct {
}

func (o *IpamVlanGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlan-groups/][%d] ipamVlanGroupsBulkDeleteNoContent ", 204)
}

func (o *IpamVlanGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
