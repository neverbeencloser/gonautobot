package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamIPAddressesBulkDeleteReader is a Reader for the IpamIPAddressesBulkDelete structure.
type IpamIPAddressesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamIPAddressesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamIPAddressesBulkDeleteNoContent creates a IpamIPAddressesBulkDeleteNoContent with default headers values
func NewIpamIPAddressesBulkDeleteNoContent() *IpamIPAddressesBulkDeleteNoContent {
	return &IpamIPAddressesBulkDeleteNoContent{}
}

/* IpamIPAddressesBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamIPAddressesBulkDeleteNoContent ipam Ip addresses bulk delete no content
*/
type IpamIPAddressesBulkDeleteNoContent struct {
}

func (o *IpamIPAddressesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-addresses/][%d] ipamIpAddressesBulkDeleteNoContent ", 204)
}

func (o *IpamIPAddressesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
