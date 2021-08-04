package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamRirsBulkDeleteReader is a Reader for the IpamRirsBulkDelete structure.
type IpamRirsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRirsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRirsBulkDeleteNoContent creates a IpamRirsBulkDeleteNoContent with default headers values
func NewIpamRirsBulkDeleteNoContent() *IpamRirsBulkDeleteNoContent {
	return &IpamRirsBulkDeleteNoContent{}
}

/* IpamRirsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamRirsBulkDeleteNoContent ipam rirs bulk delete no content
*/
type IpamRirsBulkDeleteNoContent struct {
}

func (o *IpamRirsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipamRirsBulkDeleteNoContent ", 204)
}

func (o *IpamRirsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
