package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamRouteTargetsBulkDeleteReader is a Reader for the IpamRouteTargetsBulkDelete structure.
type IpamRouteTargetsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRouteTargetsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRouteTargetsBulkDeleteNoContent creates a IpamRouteTargetsBulkDeleteNoContent with default headers values
func NewIpamRouteTargetsBulkDeleteNoContent() *IpamRouteTargetsBulkDeleteNoContent {
	return &IpamRouteTargetsBulkDeleteNoContent{}
}

/* IpamRouteTargetsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamRouteTargetsBulkDeleteNoContent ipam route targets bulk delete no content
*/
type IpamRouteTargetsBulkDeleteNoContent struct {
}

func (o *IpamRouteTargetsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/][%d] ipamRouteTargetsBulkDeleteNoContent ", 204)
}

func (o *IpamRouteTargetsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
