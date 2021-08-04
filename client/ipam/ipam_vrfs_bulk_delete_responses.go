package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVrfsBulkDeleteReader is a Reader for the IpamVrfsBulkDelete structure.
type IpamVrfsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVrfsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVrfsBulkDeleteNoContent creates a IpamVrfsBulkDeleteNoContent with default headers values
func NewIpamVrfsBulkDeleteNoContent() *IpamVrfsBulkDeleteNoContent {
	return &IpamVrfsBulkDeleteNoContent{}
}

/* IpamVrfsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamVrfsBulkDeleteNoContent ipam vrfs bulk delete no content
*/
type IpamVrfsBulkDeleteNoContent struct {
}

func (o *IpamVrfsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/][%d] ipamVrfsBulkDeleteNoContent ", 204)
}

func (o *IpamVrfsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
