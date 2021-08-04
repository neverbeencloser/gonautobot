package ipam

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamRouteTargetsDeleteReader is a Reader for the IpamRouteTargetsDelete structure.
type IpamRouteTargetsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRouteTargetsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRouteTargetsDeleteNoContent creates a IpamRouteTargetsDeleteNoContent with default headers values
func NewIpamRouteTargetsDeleteNoContent() *IpamRouteTargetsDeleteNoContent {
	return &IpamRouteTargetsDeleteNoContent{}
}

/* IpamRouteTargetsDeleteNoContent describes a response with status code 204, with default header values.

IpamRouteTargetsDeleteNoContent ipam route targets delete no content
*/
type IpamRouteTargetsDeleteNoContent struct {
}

func (o *IpamRouteTargetsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/{id}/][%d] ipamRouteTargetsDeleteNoContent ", 204)
}

func (o *IpamRouteTargetsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
