package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamServicesBulkDeleteReader is a Reader for the IpamServicesBulkDelete structure.
type IpamServicesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamServicesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamServicesBulkDeleteNoContent creates a IpamServicesBulkDeleteNoContent with default headers values
func NewIpamServicesBulkDeleteNoContent() *IpamServicesBulkDeleteNoContent {
	return &IpamServicesBulkDeleteNoContent{}
}

/* IpamServicesBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamServicesBulkDeleteNoContent ipam services bulk delete no content
*/
type IpamServicesBulkDeleteNoContent struct {
}

func (o *IpamServicesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/services/][%d] ipamServicesBulkDeleteNoContent ", 204)
}

func (o *IpamServicesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
