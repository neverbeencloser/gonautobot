package ipam

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// IpamRouteTargetsReadReader is a Reader for the IpamRouteTargetsRead structure.
type IpamRouteTargetsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRouteTargetsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRouteTargetsReadOK creates a IpamRouteTargetsReadOK with default headers values
func NewIpamRouteTargetsReadOK() *IpamRouteTargetsReadOK {
	return &IpamRouteTargetsReadOK{}
}

/* IpamRouteTargetsReadOK describes a response with status code 200, with default header values.

IpamRouteTargetsReadOK ipam route targets read o k
*/
type IpamRouteTargetsReadOK struct {
	Payload *models.RouteTarget
}

func (o *IpamRouteTargetsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/route-targets/{id}/][%d] ipamRouteTargetsReadOK  %+v", 200, o.Payload)
}
func (o *IpamRouteTargetsReadOK) GetPayload() *models.RouteTarget {
	return o.Payload
}

func (o *IpamRouteTargetsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
