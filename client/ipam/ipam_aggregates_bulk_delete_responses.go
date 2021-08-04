package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamAggregatesBulkDeleteReader is a Reader for the IpamAggregatesBulkDelete structure.
type IpamAggregatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamAggregatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamAggregatesBulkDeleteNoContent creates a IpamAggregatesBulkDeleteNoContent with default headers values
func NewIpamAggregatesBulkDeleteNoContent() *IpamAggregatesBulkDeleteNoContent {
	return &IpamAggregatesBulkDeleteNoContent{}
}

/* IpamAggregatesBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamAggregatesBulkDeleteNoContent ipam aggregates bulk delete no content
*/
type IpamAggregatesBulkDeleteNoContent struct {
}

func (o *IpamAggregatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/aggregates/][%d] ipamAggregatesBulkDeleteNoContent ", 204)
}

func (o *IpamAggregatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
