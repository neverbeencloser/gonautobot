package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamRolesDeleteReader is a Reader for the IpamRolesDelete structure.
type IpamRolesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRolesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRolesDeleteNoContent creates a IpamRolesDeleteNoContent with default headers values
func NewIpamRolesDeleteNoContent() *IpamRolesDeleteNoContent {
	return &IpamRolesDeleteNoContent{}
}

/* IpamRolesDeleteNoContent describes a response with status code 204, with default header values.

IpamRolesDeleteNoContent ipam roles delete no content
*/
type IpamRolesDeleteNoContent struct {
}

func (o *IpamRolesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/roles/{id}/][%d] ipamRolesDeleteNoContent ", 204)
}

func (o *IpamRolesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
