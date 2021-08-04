package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVlanGroupsDeleteReader is a Reader for the IpamVlanGroupsDelete structure.
type IpamVlanGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVlanGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlanGroupsDeleteNoContent creates a IpamVlanGroupsDeleteNoContent with default headers values
func NewIpamVlanGroupsDeleteNoContent() *IpamVlanGroupsDeleteNoContent {
	return &IpamVlanGroupsDeleteNoContent{}
}

/* IpamVlanGroupsDeleteNoContent describes a response with status code 204, with default header values.

IpamVlanGroupsDeleteNoContent ipam vlan groups delete no content
*/
type IpamVlanGroupsDeleteNoContent struct {
}

func (o *IpamVlanGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsDeleteNoContent ", 204)
}

func (o *IpamVlanGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
