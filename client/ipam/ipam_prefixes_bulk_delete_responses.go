package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamPrefixesBulkDeleteReader is a Reader for the IpamPrefixesBulkDelete structure.
type IpamPrefixesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamPrefixesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamPrefixesBulkDeleteNoContent creates a IpamPrefixesBulkDeleteNoContent with default headers values
func NewIpamPrefixesBulkDeleteNoContent() *IpamPrefixesBulkDeleteNoContent {
	return &IpamPrefixesBulkDeleteNoContent{}
}

/* IpamPrefixesBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamPrefixesBulkDeleteNoContent ipam prefixes bulk delete no content
*/
type IpamPrefixesBulkDeleteNoContent struct {
}

func (o *IpamPrefixesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/prefixes/][%d] ipamPrefixesBulkDeleteNoContent ", 204)
}

func (o *IpamPrefixesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
